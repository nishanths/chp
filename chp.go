// Package chp implements common channel patterns.
// The channel element types are defined using type parameters.
package chp

import (
	"sync"
)

// first returns the last value received from any of the input channels.
func last[T any](cs ...chan T) T {
	done := make(chan struct{})
	defer close(done)
	var last T
	for v := range merge(done, 0, cs...) {
		last = v
	}
	return last
}

// first returns the first value received from any of the input channels.
func first[T any](cs ...chan T) T {
	done := make(chan struct{})
	defer close(done)
	return <-merge(done, 0, cs...)
}

// merge multiplexes values from multiple input channels into a single output
// channel. (merge, internally, starts one goroutine per input channel.) merge
// closes the output channel and releases resources either when all input
// channels are closed or when done is closed.
func merge[T any](done <-chan struct{}, buffer int, cs ...chan T) <-chan T {
	out := make(chan T, buffer)
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

// collect receives values from a channel and collects them in a slice. The
// slice is returned when the channel is closed.
func collect[T any](c chan T) []T {
	var out []T
	for v := range c {
		out = append(out, v)
	}
	return out
}

// partition sends each value received from the channel to either the t channel
// or the f channel, based on whether the predicate function, applied to the
// value, returns true or false, respectively.
func partition[T any](c <-chan T, predicate func(T) bool) (t, f <-chan T) {
	tt := make(chan T, cap(c))
	ff := make(chan T, cap(c))
	go func() {
		defer close(ff)
		defer close(tt)
		for v := range c {
			if predicate(v) {
				tt <- v
			} else {
				ff <- v
			}
		}
	}()
	return tt, ff
}

// mapchan sends each value received from the input channel
// to the input channel after applying the mapping function f.
// The output channel is closed when the input channel is closed.
func mapchan[T, U any](c <-chan T, f func(T) U) <-chan U {
	out := make(chan U, cap(c))
	go func() {
		defer close(out)
		for v := range c {
			out <- f(v)
		}
	}()
	return out
}

// choose sends to the output channel only those values
// from the input channel that satisfy the predicate.
// The output channel is closed when the input channel is closed.
func choose[T any](c <-chan T, predicate func(T) bool) <-chan T {
	out := make(chan T, cap(c))
	go func() {
		defer close(out)
		for v := range c {
			if predicate(v) {
				out <- v
			}
		}
	}()
	return out
}

// drop sends to the output channel only those values
// from the input channel that do not satisfy the predicate.
// The output channel is closed when the input channel is closed.
func drop[T any](c <-chan T, predicate func(T) bool) <-chan T {
	reverse := func(x T) bool { return !predicate(x) }
	return choose(c, reverse)
}
