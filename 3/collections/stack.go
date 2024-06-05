package collections

import "sync"

type Stack struct {
	new   bool
	mx    *sync.RWMutex
	slice []any
}

//  Нужен ли мне конструктор?
// 	Наверное да, как минимум для инициализации мьютекса
// 	Если я захочу ли я сделать так, чтобы тип слайса заранее был определен, чтобы я в нем мог хранить только определенный тип?
// 		- reflection - чтобы при создании стэка мы указывали тип данных
// 		- обертка над Stack New\NewInt\NewString - не используем рефлект, что плюс, но не писать же на каждый тип?
// 		Пока это под вопросом, сомнения:
//	Обертки над каждым типом писать - как будто не правильно, но и рефлект - нигде не рекоммендуют использовать
//  Интерфейс? Можно выделить методы: получить последнее \ первое значение \ удалить значение \ сократить слайс
//	Пока оставлю это под вопросом

func NewStack() *Stack {
	mx := sync.RWMutex{}
	return &Stack{new: true, mx: &mx}
}

func (s *Stack) wasnew() {
	if !s.new {
		panic("was initialized without NewStack()")
	}
}

func (s *Stack) Clear() {
	s.wasnew()
	s.mx.Lock()
	defer s.mx.Unlock()
	s.slice = nil
}

func (s *Stack) Contains(value any) bool {
	s.wasnew()
	s.mx.RLock()
	defer s.mx.RUnlock()
	if len(s.slice) <= 0 {
		return false
	}
	for _, v := range s.slice {
		if v == value {
			return true
		}
	}
	return false
}

func (s *Stack) Push(value any) {
	s.wasnew()
	s.mx.Lock()
	defer s.mx.Unlock()
	s.slice = append(s.slice, value)
}

func (s *Stack) Pop() any {
	s.wasnew()
	s.mx.Lock()
	defer s.mx.Unlock()
	if len(s.slice) <= 0 {
		return nil
	}
	res := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return res
}

func (s *Stack) Peek() any {
	s.wasnew()
	s.mx.RLock()
	defer s.mx.RUnlock()
	if len(s.slice) <= 0 {
		return nil
	}
	return s.slice[len(s.slice)-1]
}

func (s *Stack) TryPop() (any, bool) {
	s.wasnew()
	s.mx.Lock()
	defer s.mx.Unlock()
	if len(s.slice) <= 0 {
		return nil, false
	}
	res := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return res, true
}

func (s *Stack) TryPeek() (any, bool) {
	s.wasnew()
	s.mx.RLock()
	defer s.mx.RUnlock()
	if len(s.slice) <= 0 {
		return nil, false
	}
	res := s.slice[len(s.slice)-1]
	return res, true
}
