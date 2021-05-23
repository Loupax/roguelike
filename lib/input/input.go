package input

import (
	"bufio"
	"context"
	"io"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func SetupTerminal(fd int) (func(), error) {
	old, err := terminal.MakeRaw(fd)
	if err != nil {
		return func() {}, err
	}
	return func() {
		terminal.Restore(fd, old)
	}, nil
}

type keyboard struct {
	ch     chan rune
	err    error
	closed bool
}

func (k *keyboard) done() {
	k.closed = true
	k.ch <- 0
	close(k.ch)
}

func NewKeyboard(ctx context.Context) *keyboard {
	ch := make(chan rune)
	ioCh := make(chan rune)
	errCh := make(chan error)
	k := keyboard{ch: ch}
	// Why the two goroutines?
	//
	// Because reading from stdin
	// is blocking, we cannot cancel/timeout a goroutine until a keypress is handled
	// This defeats the entire purpose of the cancelation context.
	//
	// The solution is to split the reading from stdin and sending it's contents
	// in two separate goroutines. That way we can have a single goroutine that
	// only fires when we receive a cancellation signal OR when we receive
	// input from another channel, and another goroutine that just writes
	// to the afformation channel.
	go func(ioCh, ch chan rune, errCh chan error) {
		for {
			select {
			case <-ctx.Done():
				k.done()
				return
			case b := <-ioCh:
				ch <- b
			case err := <-errCh:
				k.err = err
				k.done()
				return
			}
		}
	}(ioCh, ch, errCh)

	go func(ioCh chan rune, errCh chan error) {
		stdinReader := bufio.NewReaderSize(os.Stdin, 1)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				b, _, err := stdinReader.ReadRune()
				if err != nil {
					errCh <- err
					return
				} else {
					ioCh <- b
				}
			}
		}
	}(ioCh, errCh)
	return &k
}

func (k *keyboard) KeyPress(b *rune) error {
	if k.err != nil {
		return k.err
	}
	if k.closed {
		return io.EOF
	}
	*b = <-k.ch
	return nil
}
