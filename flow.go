package errflow

// Flow is the base error handling interface.
type Flow interface {
	// Err returns any error if one occurred
	Err() error
	// Fail ends this flow with the error
	Fail(error)
}

// Linkable is a composition interface for Flow
type Linkable interface {
	// Link returns the base error this element
	Link() *error
	// LinkTo mus replace the base error with the new error
	LinkTo(*error)
}

// LinkFlow links new to old making old fail if new fails and vice versa
func LinkFlow[T Linkable](old Linkable, new T) T {
	old.LinkTo(new.Link())
	return new
}

// LinkableFlow combines Flow and Linkable
type LinkableFlow interface {
	Linkable
	Flow
}

type errChain struct {
	err *error
}

func errChainOf(flow Linkable) errChain {
	return errChain{err: flow.Link()}
}

func errChainOfErr(err error) errChain {
	return errChain{err: &err}
}

func emptyChain() errChain {
	return errChain{
		err: new(error),
	}
}

func (chain *errChain) Err() error {
	return *chain.err
}

func (chain *errChain) Fail(err error) {
	if chain.err != nil {
		chain.err = &err
	}
}

func (chain *errChain) Link() *error {
	return chain.err
}

func (chain *errChain) LinkTo(err *error) {
	if *chain.err != nil {
		*err = *chain.err
	}
	chain.err = err
}

// Do executes doFunc which could fail.
// Returns doFunc result if successful or the result of or.
// Respects a already failed flow and fails the flow correclty.
func Do[T any](doFunc func() (T, error), flow Flow, or func(error) T) T {
	err := flow.Err()
	if err != nil {
		return or(err)
	}
	res, err := doFunc()
	if err != nil {
		flow.Fail(err)
		return or(err)
	}
	return res
}

func pass[T Flow](doFunc func() error, flow T) T {
	if flow.Err() != nil {
		return flow
	}
	err := doFunc()
	if err != nil {
		flow.Fail(err)
	}
	return flow
}

type linkedFlow struct {
	err *error
}

func empty[T any](_ error) (t T) { return }
