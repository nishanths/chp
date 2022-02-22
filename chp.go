// Package chp implements common channel patterns. The channel element types are
// defined using type parameters.
package chp

import (
	"sync"
)

// First returns the first value received from any of the input channels.
func First[T any](cs ...chan T) T {
	done := make(chan struct{})
	defer close(done) // release resources
	return <-Merge(done, cs...)
}

// Merge multiplexes values from multiple input channels into a single output
// channel. The output channel is closed when all input channels are closed, or
// when the done channel is closed. If the output stream is no longer needed,
// close the done channel to release resources used by Merge.
func Merge[T any](done <-chan struct{}, cs ...chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	go func() {
		wg.Wait()
		close(out)
	}()

	for _, c := range cs {
		go func(c <-chan T) {
			defer wg.Done()
			for {
				select {
				case v, ok := <-c:
					if !ok {
						return
					}
					out <- v
				case <-done:
					return
				}
			}
		}(c)
	}

	return out
}

// Collect receives values from a channel and collects them in a slice. The
// slice is returned when the channel is closed.
func Collect[T any](c chan T) []T {
	var out []T
	for v := range c {
		out = append(out, v)
	}
	return out
}
