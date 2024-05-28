package point

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type point struct {
	longtitude float64
	latitude   float64
}

func New(longtitude float64, latitude float64) point {
	return point{longtitude: longtitude, latitude: latitude}
}

func NewPointString(s string) (point, error) {
	coords := strings.Split(s, ";")
	if len(coords) != 2 {
		return point{}, errors.New("incorrect coords")
	}
	long, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return point{}, err
	}
	lat, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return point{}, err
	}
	return point{long, lat}, nil
}

func GetDistance(p1 point, p2 point) float64 {
	return math.Sqrt(math.Pow((p1.longtitude-p2.longtitude), 2) + math.Pow((p1.latitude-p2.latitude), 2))
}

func PointInCircle(p1 point, p2 point, radius float64) bool {
	return (math.Pow((p1.longtitude-p2.longtitude), 2) + math.Pow((p1.latitude-p1.latitude), 2)) <= math.Pow(radius, 2)
}
