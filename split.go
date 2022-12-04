package errflow

// Splitted is a helper for more fluent code when working with return values.
//
// Example:
//   myFlow.DoSomething(). // returns a splitted flow
//     Run(func(value V) { doSomethingWithV(value) }). // returns myFlow
//     BackToMyFlow()
type Splitted[F any, S Flow] struct {
	f F
	s S
}

// SplitOf splits the value f from the flow s
func SplitOf[F any, S Flow](f F, s S) Splitted[F, S] {
	return Splitted[F, S]{f, s}
}

// Run executes run with the splitted value and returns the originial flow.
// Is a noop when the flow already failed.
func (splitted Splitted[F, S]) Run(run func(F)) S {
	if splitted.s.Err() != nil {
		return splitted.s
	}
	run(splitted.f)
	return splitted.s
}

// RunErr runs run and might fail the flow with it if an error is encountered.
// Is a noop when the flow already failed.
func (splitted Splitted[F, S]) RunErr(run func(F) error) S {
	if splitted.s.Err() != nil {
		return splitted.s
	}
	err := run(splitted.f)
	if err != nil {
		splitted.s.Fail(err)
	}
	return splitted.s
}

// Ignore returns the flow.
func (splitted Splitted[F, S]) Ignore() S {
	return splitted.s
}

// Collapse returns the splitted value.
func (splitted Splitted[F, S]) Collapse() F {
	return splitted.f
}
