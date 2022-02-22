package chp

// ReceiveOnly provides a way to make a bidrectional channel receive-only from
// the type system's perspective, without having to write the channel element type
// explicitly in source code.
//
// Note that this function cannot prevent code that already has a reference to c
// as a bidirectional channel from using it bidirectionally.
func ReceiveOnly[T any](c chan T) <-chan T {
	var d <-chan T
	d = c
	return d
}

// SendOnly provides a way to make a bidrectional channel send-only from the
// type system's perspective, without having to write the channel element type
// explicitly in source code.
//
// Note that this function cannot prevent code that already has a reference to c
// as a bidirectional channel from using it bidirectionally.
func SendOnly[T any](c chan T) chan<- T {
	var d chan<- T
	d = c
	return d
}
