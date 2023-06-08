# chp

Common channel patterns implemented using type parameters.

The functions in this package are intentionally not exported. Copy and maintain
any necessary functions manually.

## Docs

Output of `go doc -u -all`:

```
package chp // import "github.com/nishanths/chp"

Package chp implements common channel patterns. The channel element types are
defined using type parameters.

FUNCTIONS

func choose[T any](c <-chan T, predicate func(T) bool) <-chan T
    choose sends to the output channel only those values from the input channel
    that satisfy the predicate. The output channel is closed when the input
    channel is closed.

func collect[T any](c chan T) []T
    collect receives values from a channel and collects them in a slice.
    The slice is returned when the channel is closed.

func drop[T any](c <-chan T, predicate func(T) bool) <-chan T
    drop sends to the output channel only those values from the input channel
    that do not satisfy the predicate. The output channel is closed when the
    input channel is closed.

func first[T any](cs ...chan T) T
    first returns the first value received from any of the input channels.

func last[T any](cs ...chan T) T
    first returns the last value received from any of the input channels.

func mapchan[T, U any](c <-chan T, f func(T) U) <-chan U
    mapchan sends each value received from the input channel to the input
    channel after applying the mapping function f. The output channel is closed
    when the input channel is closed.

func merge[T any](done <-chan struct{}, buffer int, cs ...chan T) <-chan T
    merge multiplexes values from multiple input channels into a single output
    channel. (merge, internally, starts one goroutine per input channel.) merge
    closes the output channel and releases resources either when all input
    channels are closed or when done is closed.

func partition[T any](c <-chan T, predicate func(T) bool) (t, f <-chan T)
    partition sends each value received from the channel to either the t channel
    or the f channel, based on whether the predicate function, applied to the
    value, returns true or false, respectively.

func receiveOnly[T any](c chan T) (r <-chan T)
    receiveOnly provides a way to make a bidrectional channel receive-only
    to the type system, without having to specify the channel element type
    explicitly in source code.

func sendOnly[T any](c chan T) (s chan<- T)
    sendOnly provides a way to make a bidrectional channel send-only to the type
    system, without having to specify the channel element type explicitly in
    source code.
```
