package collections

type Dueue struct {
	Collection
}

func NewDeque() *Dueue {
	c := New()
	return &Dueue{*c}
}
