package main

type LimitSupport struct {
	*AbstractSupport
	limit int
}

func NewLimitSupport(name string, limit int) *LimitSupport {
	return &LimitSupport{
		AbstractSupport: NewAbstractSupport(name),
		limit:           limit,
	}
}

func (s *LimitSupport) Resolve(trouble *Trouble) bool {
	return trouble.GetNumber() < s.limit
}
