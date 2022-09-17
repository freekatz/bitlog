package event

type CollectorEventFilter struct {
}

func FromStrFilter(filterStr string) *CollectorEventFilter {
	return &CollectorEventFilter{}
}

func (f *CollectorEventFilter) FilterOut(ce CollectorEvent) bool {
	// TODO
	return false
}
