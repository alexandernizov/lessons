package set_test

import (
	"testing"

	"github.com/alexandenrizov/lessons/4/set"
	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	s1 := set.New()
	s1.Add(1)
	s1.Add(2)
	s2 := set.New()
	s2.Add(2)
	s2.Add(3)
	expected := set.New()
	expected.Add(1)
	expected.Add(2)
	expected.Add(3)
	assert.Equal(t, expected, s1.Union(s2), "Incorrect result")
}

func TestIntersection(t *testing.T) {
	s1 := set.New()
	s1.Add(1)
	s1.Add(2)
	s2 := set.New()
	s2.Add(2)
	s2.Add(3)
	expected := set.New()
	expected.Add(2)
	assert.Equal(t, expected, s1.Intersection(s2), "Incorrect result")
}

func TestSubstraction(t *testing.T) {
	s1 := set.New()
	s1.Add(1)
	s1.Add(2)
	s2 := set.New()
	s2.Add(2)
	s2.Add(3)
	expected := set.New()
	expected.Add(1)
	assert.Equal(t, expected, s1.Substraction(s2), "Incorrect result")
}

func TestRemove(t *testing.T) {
	s := set.New()
	s.Add(1)
	s.Remove(1)
	_, ok := s.M[1]
	assert.Equal(t, false, ok, "Incorrect result")
}
