package point

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	longtitude float64
	latitude   float64
}

func New(longtitude float64, latitude float64) Point {
	return Point{longtitude: longtitude, latitude: latitude}
}

func NewPointString(s string) (Point, error) {
	s = strings.ReplaceAll(s, " ", "")
	coords := strings.Split(s, ",")
	if len(coords) != 2 {
		return Point{}, errors.New("incorrect coords")
	}
	long, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return Point{}, err
	}
	lat, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return Point{}, err
	}
	return Point{long, lat}, nil
}

func GetDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow((p1.longtitude-p2.longtitude), 2) + math.Pow((p1.latitude-p2.latitude), 2))
}

func PointInCircle(p1 Point, p2 Point, radius float64) bool {
	return (math.Pow((p1.longtitude-p2.longtitude), 2) + math.Pow((p1.latitude-p1.latitude), 2)) <= math.Pow(radius, 2)
}
