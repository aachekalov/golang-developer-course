package step1

import (
	"math"
)

const eps = 0.00001

func IsTheSameLocation(lat1, lon1, lat2, lon2 float32) bool {
	return math.Abs(float64(lat1-lat2)) < eps && math.Abs(float64(lon1-lon2)) < eps
}
