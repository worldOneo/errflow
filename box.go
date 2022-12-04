package errflow

/*
type Box[T any] struct {
	sealed  bool
	present bool
	value   T
	errs    errChain
}

func EmptyBox[T any](flow Linkable) Box[T] {
	return Box[T]{
		sealed:  false,
		present: false,
		errs:    errChainOf(flow),
	}
}

func FailedBox[T any](err error) Box[T] {
	return Box[T]{
		sealed: true,
		errs:   errChainOfErr(err),
	}
}

func BoxOf[T any](fn func() (T, error), flow Linkable) Box[T] {
	chain := errChainOf(flow)
	v, err := fn()
	if err != nil {
		chain.Fail(err)
	}
	return Box[T]{
		sealed:  true,
		present: err != nil,
		value:   v,
		errs:    chain,
	}
}

func (box *Box[T]) Put(fn func() T) *Box[T] {
	if box.sealed {
		return box
	}
	if box.Err() != nil {
		return box
	}
	box.present = true
	box.sealed = true
	box.value = fn()
	return box
}

func (box *Box[T]) Fill(fn func() (T, error)) *Box[T] {
	if box.sealed {
		return box
	}
	if box.Err() != nil {
		return box
	}
	v, err := fn()
	if err != nil {
		box.errs.Fail(err)
		return box
	}
	box.sealed = true
	box.present = true
	box.value = v
	return box
}

func (box *Box[T]) Get() (T, bool) {
	return box.value, box.present
}

func (box *Box[T]) Err() error {
	return box.errs.Err()
}

func (box *Box[T]) Fail(err error) {
	box.errs.Fail(err)
}
*/
