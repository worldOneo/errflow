package errflow

type Splitted[F any, S Flow] struct {
	f F
	s S
}

func SplitOf[F any, S Flow](f F, s S) Splitted[F, S] {
	return Splitted[F, S]{f, s}
}

func (splitted Splitted[F, S]) Run(run func(F)) S {
	if splitted.s.Err() != nil {
		return splitted.s
	}
	run(splitted.f)
	return splitted.s
}

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

func (splitted Splitted[F, S]) Ignore() S {
	return splitted.s
}

func (splitted Splitted[F, S]) Collapse() F {
	return splitted.f
}
