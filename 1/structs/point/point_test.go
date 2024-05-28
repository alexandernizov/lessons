package point

import (
	"testing"
)

func TestNew(t *testing.T) {
	testTable := []struct {
		coords   []float64
		expected Point
	}{
		{
			coords:   []float64{1, 2},
			expected: Point{1, 2},
		},
		{
			coords:   []float64{1.0, 2.0},
			expected: Point{1, 2},
		},
		{
			coords:   []float64{-1, -2},
			expected: Point{-1, -2},
		},
	}
	for _, testCase := range testTable {
		p := New(testCase.coords[0], testCase.coords[1])
		if p != testCase.expected {
			t.Errorf("Incorrect result. Exptect: %v, got %v", testCase.expected, p)
		}
	}
}

func TestPointString(t *testing.T) {
	testTable := []struct {
		coords   string
		expected Point
	}{
		{
			coords:   "1,2",
			expected: Point{1, 2},
		},
		{
			coords:   "1, 2",
			expected: Point{1, 2},
		},
		{
			coords:   "1.1,2.1",
			expected: Point{1.1, 2.1},
		},
		{
			coords:   "1.1, 2.1",
			expected: Point{1.1, 2.1},
		},
		{
			coords:   "-1.1, -2.1",
			expected: Point{-1.1, -2.1},
		},
	}
	for _, testCase := range testTable {
		p, err := NewPointString(testCase.coords)
		if err != nil {
			t.Error("%w", err)
		}
		if p != testCase.expected {
			t.Errorf("Incorrect result. Expected: %v, got %v", testCase.expected, p)
		}
	}

}

func TestGetDistance(t *testing.T) {
	testTable := []struct {
		points   []Point
		expected float64
	}{
		{
			points:   []Point{{0, 0}, {0, 0}},
			expected: 0,
		},
		{
			points:   []Point{{0, 0}, {3, 4}},
			expected: 5,
		},
	}
	for _, testCase := range testTable {
		p := GetDistance(testCase.points[0], testCase.points[1])
		if p != testCase.expected {
			t.Errorf("Incorrect result. Expected: %v, got %v", testCase.expected, p)
		}
	}
}

func TestPointInCircle(t *testing.T) {
	testTable := []struct {
		points   []Point
		radius   float64
		expected bool
	}{
		{
			points:   []Point{{0, 0}, {0, 0}},
			radius:   0,
			expected: true,
		},
		{
			points:   []Point{{0, 0}, {0, 0}},
			radius:   1,
			expected: true,
		},
		{
			points:   []Point{{0, 0}, {1, 1}},
			radius:   0,
			expected: false,
		},
		{
			points:   []Point{{0, 0}, {1, 1}},
			radius:   2,
			expected: true,
		},
	}
	for _, testCase := range testTable {
		p := PointInCircle(testCase.points[0], testCase.points[1], testCase.radius)
		if p != testCase.expected {
			t.Errorf("Incorrect result. Expected: %v, got %v", testCase.expected, p)
		}
	}
}
