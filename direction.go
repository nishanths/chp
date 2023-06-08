package chp

// receiveOnly provides a way to make a bidrectional channel receive-only to the
// type system, without having to specify the channel element type explicitly
// in source code.
func receiveOnly[T any](c chan T) (r <-chan T) {
	r = c
	return
}

// sendOnly provides a way to make a bidrectional channel send-only to the
// type system, without having to specify the channel element type explicitly in
// source code.
func sendOnly[T any](c chan T) (s chan<- T) {
	s = c
	return
}
