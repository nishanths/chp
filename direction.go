package chp

// ReceiveOnly provides a way to make a bidrectional channel receive-only from
// the type system's perspective, without having to specify the channel's
// element type explicitly in source code.
func ReceiveOnly[T any](c chan T) (r <-chan T) {
	r = c
	return
}

// SendOnly provides a way to make a bidrectional channel send-only from the
// type system's perspective, without having to specify the channel's element
// type explicitly in source code.
func SendOnly[T any](c chan T) (s chan<- T) {
	s = c
	return
}
