package set

type Set struct {
	M map[any]struct{}
}

//"Пересечение двух множеств", "Объединение двух множеств", "Вычитание множества"

func New() *Set {
	return &Set{M: make(map[any]struct{})}
}

func (s *Set) Add(value any) {
	s.M[value] = struct{}{}
}

func (s *Set) Remove(value any) {
	delete(s.M, value)
}

func (s *Set) Union(s2 *Set) *Set {
	res := New()
	for k := range s.M {
		res.M[k] = struct{}{}
	}
	for k := range s2.M {
		res.M[k] = struct{}{}
	}
	return res
}

func (s *Set) Intersection(s2 *Set) *Set {
	res := New()
	for k := range s.M {
		_, ok := s2.M[k]
		if ok {
			res.M[k] = struct{}{}
		}
	}
	return res
}

func (s *Set) Substraction(s2 *Set) *Set {
	res := New()
	for k := range s.M {
		_, ok := s2.M[k]
		if ok {
			continue
		}
		res.M[k] = struct{}{}
	}
	return res
}
