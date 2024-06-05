package collections

//Скорее всего это идеологически неправильно и надо сделать либо отдельные структуры, либо интерфейс, но сделал так

type Deque struct {
	Queue
}

func NewDeque() *Deque {
	q := NewQueue()
	return &Deque{Queue: *q}
}

func (d *Deque) PushFront(value any) {
	d.wasnew()
	d.mx.Lock()
	defer d.mx.Unlock()
	d.slice = append(d.slice, value)
	copy(d.slice[1:], d.slice)
	d.slice[0] = value
}

func (d *Deque) PopBack() any {
	d.wasnew()
	d.mx.Lock()
	defer d.mx.Unlock()
	if len(d.slice) <= 0 {
		return nil
	}
	res := d.slice[len(d.slice)-1]
	d.slice = d.slice[:len(d.slice)-1]
	return res
}
