// Package chp implements common channel patterns.
//
// The channel element types are defined using type parameters.
package chp

import "sync"

// First returns the first value received from any of the channels.
func First[T any](cs ...chan T) T {
	done := make(chan struct{})
	defer close(done) // release resources
	return <-FanIn(done, cs...)
}

// FanIn returns an output channel that merges the values from all of the input
// channels.
//
// Value order is maintained in the output stream within each input channel, but
// not across input channels. The output channel is closed when all of the input
// channels are closed or when the done channel is closed. If the output stream
// is no longer needed, close the done channel to free resources used by
// FanIn.
func FanIn[T any](done <-chan struct{}, cs ...chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(cs))

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

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
