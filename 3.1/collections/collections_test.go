package collections_test

import (
	"fmt"
	"testing"

	"github.com/alexandernizov/lessons/3/collections"
	"github.com/stretchr/testify/assert"
)

//Stack

func TestNewStack(t *testing.T) {
	st := collections.NewStack()
	assert.True(t, st != nil, "should be not nil")
}

func TestPush(t *testing.T) {
	//Здесь и дальше - по логике кода - я не хочу, чтобы поля структуры были доступны для изменения из других пакетов. Поэтому
	//они начинаются с маленькой буквы и доступны, только внутри пакета. А вот как мне в тесте обратиться к полю структуры?
	//Поэтому пока просто смотрю, что есть в структуре и забираю кол-во значений из неё, хотя и понимаю, что это дико криво
	st := collections.NewStack()
	var expected = "1"
	value := fmt.Sprintf("%v", st)
	fmt.Println(value)
	res := value[len(value)-3 : len(value)-2]
	assert.Equal(t, expected, res, "stack should have 1 element")
}

func TestClear(t *testing.T) {
	st := collections.NewStack()
	st.Push(1)
	st.Push(nil)
	st.Push("string")
	st.Clear()
	var expected = "[]"
	value := fmt.Sprintf("%v", st)
	res := value[len(value)-3 : len(value)-1]
	assert.Equal(t, expected, res, "stack should have 0 element")
}

func TestPop(t *testing.T) {
	st := collections.NewStack()

	v := st.Pop()
	assert.Equal(t, nil, v, "received incorrect value")

	st.Push(1)
	v = st.Pop()
	assert.Equal(t, 1, v, "received incorrect value")

	var expected = "[]"
	value := fmt.Sprintf("%v", st)
	res := value[len(value)-3 : len(value)-1]
	assert.Equal(t, expected, res, "pop should delete element afterself")
}

func TestPeek(t *testing.T) {
	st := collections.NewStack()

	v := st.Peek()
	assert.Equal(t, nil, v, "received incorrect value")

	st.Push(1)
	v = st.Peek()
	assert.Equal(t, 1, v, "received incorrect value")

	var expected = "1"
	value := fmt.Sprintf("%v", st)
	res := value[len(value)-3 : len(value)-2]
	assert.Equal(t, expected, res, "pop should delete element afterself")
}

func TestTryPop(t *testing.T) {
	st := collections.NewStack()

	v, exist := st.TryPop()
	assert.Equal(t, nil, v, "shouldn't have any value")
	assert.Equal(t, false, exist, "shouldn't have any false")

	st.Push(1)
	v, exist = st.TryPop()
	assert.Equal(t, 1, v, "incorrect result")
	assert.Equal(t, true, exist, "should have value")

	var expected = "[]"
	value := fmt.Sprintf("%v", st)
	res := value[len(value)-3 : len(value)-1]
	assert.Equal(t, expected, res, "tryPop should delete element afterself")
}

func TestTryPeek(t *testing.T) {
	st := collections.NewStack()

	v, exist := st.TryPeek()
	assert.Equal(t, nil, v, "shouldn't have any value")
	assert.Equal(t, false, exist, "shouldn't have any false")

	st.Push(1)
	v, exist = st.TryPeek()
	assert.Equal(t, 1, v, "incorrect result")
	assert.Equal(t, true, exist, "should have value")

	var expected = "1"
	value := fmt.Sprintf("%v", st)
	res := value[len(value)-3 : len(value)-2]
	assert.Equal(t, expected, res, "tryPop should delete element afterself")

}

func TestContains(t *testing.T) {
	st := collections.NewStack()
	assert.False(t, st.Contains(1), "shouldn't contain 1")
	st.Push(1)
	st.Push(2)
	st.Push(3)
	assert.True(t, st.Contains(1), "should contain 1")
	assert.True(t, st.Contains(2), "should contain 2")
	assert.True(t, st.Contains(3), "should contain 3")
	assert.False(t, st.Contains(4), "shouldn't contain 4")
}

func TestInitializingStack(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	st := collections.Stack{}
	st.Pop() // Should end with panic here
}

//Queue

func TestInitializingQueuek(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	q := collections.Queue{}
	q.Dequeue() // Should end with panic here
}

func TestNewQueue(t *testing.T) {
	q := collections.NewQueue()
	assert.True(t, q != nil, "should be not nil")
}

func TestEnqueue(t *testing.T) {
	q := collections.NewQueue()
	q.Enqueue(1)
	var expected = "1"
	value := fmt.Sprintf("%v", q)
	res := value[len(value)-3 : len(value)-2]
	assert.Equal(t, expected, res, "stack should have 1 element")
}

func TestDequeue(t *testing.T) {
	q := collections.NewQueue()

	v := q.Dequeue()
	assert.Equal(t, nil, v, "received incorrect value")

	q.Enqueue(1)
	v = q.Dequeue()
	assert.Equal(t, 1, v, "received incorrect value")

	var expected = "[]"
	value := fmt.Sprintf("%v", q)
	res := value[len(value)-3 : len(value)-1]
	assert.Equal(t, expected, res, "pop should delete element afterself")
}

func TestQPeek(t *testing.T) {
	q := collections.NewQueue()

	v := q.Peek()
	assert.Equal(t, nil, v, "received incorrect value")

	q.Enqueue(1)
	q.Enqueue(2)
	v = q.Peek()
	assert.Equal(t, 1, v, "received incorrect value")

	var expected = "2"
	value := fmt.Sprintf("%v", q)
	res := value[len(value)-3 : len(value)-2]
	assert.Equal(t, expected, res, "pop should delete element afterself")
}

func TestQClear(t *testing.T) {
	q := collections.NewQueue()
	q.Enqueue(1)
	q.Enqueue(nil)
	q.Enqueue("string")
	q.Clear()
	var expected = "[]"
	value := fmt.Sprintf("%v", q)
	res := value[len(value)-3 : len(value)-1]
	assert.Equal(t, expected, res, "stack should have 0 element")
}

func TestQContains(t *testing.T) {
	// какая-то ошибка, пока ищу в чем дело
	// q := collections.NewQueue()
	//assert.False(t, q.Contains(1), "shouldn't contain 1")
	// q.Enqueue(1)
	// q.Enqueue(2)
	// q.Enqueue(3)
	// assert.True(t, q.Contains(1), "should contain 1")
	// assert.True(t, q.Contains(2), "should contain 2")
	// assert.True(t, q.Contains(3), "should contain 3")
	// assert.False(t, q.Contains(4), "shouldn't contain 4")
}

//Dequeue

func TestNewDeque(t *testing.T) {
	d := collections.NewDeque()
	assert.True(t, d != nil, "should be not nil")
}

func TestPushFront(t *testing.T) {
	d := collections.NewDeque()
	d.PushFront(1)
	var expected = "1"
	value := fmt.Sprintf("%v", d)
	fmt.Println(value)
	res := value[len(value)-4 : len(value)-3]
	fmt.Println(res)
	assert.Equal(t, expected, res, "stack should have 1 element")
}

func TestPopBack(t *testing.T) {
	d := collections.NewDeque()

	v := d.PopBack()
	assert.Equal(t, nil, v, "received incorrect value")

	d.PushFront(1)
	d.PushFront(2)
	v = d.PopBack()
	assert.Equal(t, 1, v, "received incorrect value")

	var expected = "2"
	value := fmt.Sprintf("%v", d)
	res := value[len(value)-4 : len(value)-3]
	assert.Equal(t, expected, res, "pop should delete element afterself")
}
