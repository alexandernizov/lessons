package polygon

import (
	"errors"
	"strings"

	"github.com/alexandernizov/lessons/1/structs/point"
)

type Polygon struct {
	points []point.Point
}

func New(points ...point.Point) (Polygon, error) {
	if len(points) < 3 {
		return Polygon{}, errors.New("polygon should include minimum 3 points")
	}
	return Polygon{points: points}, nil
}

func NewPolygonString(s string) (Polygon, error) {
	s = strings.ReplaceAll(s, " ", "")
	pointsString := strings.Split(s, ";")
	if len(pointsString) < 3 {
		return Polygon{}, errors.New("polygon should include minimum 3 points")
	}
	var points []point.Point
	for _, coords := range pointsString {
		p, err := point.NewPointString(coords)
		if err != nil {
			return Polygon{}, err
		}
		points = append(points, p)
	}
	return Polygon{points: points}, nil
}
