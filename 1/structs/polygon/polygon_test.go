package polygon

import (
	"reflect"
	"testing"

	"github.com/alexandernizov/lessons/1/structs/point"
)

func TestNew(t *testing.T) {
	testTable := []struct {
		points   []point.Point
		expected polygon
		err      string
	}{
		{
			points:   []point.Point{},
			expected: polygon{},
			err:      "polygon should include minimum 3 points",
		},
		{
			points: []point.Point{
				point.New(1, 2), point.New(3, 4), point.New(5, 6), point.New(7, 8),
			},
			expected: polygon{
				points: []point.Point{point.New(1, 2), point.New(3, 4), point.New(5, 6), point.New(7, 8)},
			},
			err: "polygon should include minimum 3 points",
		},
	}

	for _, testCase := range testTable {
		p, err := New(testCase.points...)
		if err != nil && (err.Error() != testCase.err) {
			t.Errorf("%v", err)
		}
		if !reflect.DeepEqual(p, testCase.expected) {
			t.Errorf("Incorrect result. Exptect: %v, got %v", testCase.expected, p)
		}
	}
}

func TestNewPolygonString(t *testing.T) {
	testTable := []struct {
		input    string
		expected polygon
		err      string
	}{
		{
			input:    "",
			expected: polygon{},
			err:      "polygon should include minimum 3 points",
		},
		{
			input: "1,2; 3,4; 5,6; 7,8",
			expected: polygon{
				points: []point.Point{point.New(1, 2), point.New(3, 4), point.New(5, 6), point.New(7, 8)},
			},
			err: "",
		},
	}

	for _, testCase := range testTable {
		p, err := NewPolygonString(testCase.input)
		if err != nil && (err.Error() != testCase.err) {
			t.Errorf("%v", err)
		}
		if !reflect.DeepEqual(p, testCase.expected) {
			t.Errorf("Incorrect result. Exptect: %v, got %v", testCase.expected, p)
		}
	}
}
