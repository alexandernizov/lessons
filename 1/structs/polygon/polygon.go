package polygon

import (
	"errors"
	"strings"

	"github.com/alexandernizov/lessons/1/structs/point"
)

type polygon struct {
	points []point.Point
}

func New(points ...point.Point) (polygon, error) {
	if len(points) < 3 {
		return polygon{}, errors.New("polygon should include minimum 3 points")
	}
	return polygon{points: points}, nil
}

func NewPolygonString(s string) (polygon, error) {
	s = strings.ReplaceAll(s, " ", "")
	pointsString := strings.Split(s, ";")
	if len(pointsString) < 3 {
		return polygon{}, errors.New("polygon should include minimum 3 points")
	}
	var points []point.Point
	for _, coords := range pointsString {
		p, err := point.NewPointString(coords)
		if err != nil {
			return polygon{}, err
		}
		points = append(points, p)
	}
	return polygon{points: points}, nil
}
