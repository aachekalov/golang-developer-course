package step1

import (
	"math"
)

func GetDistanceFromLatLonInKm(lat1, lon1, lat2, lon2 float32) float64 {
	var R = 6371.0 // Radius of the earth in km
	var dLat = deg2rad(lat2 - lat1)
	var dLon = deg2rad(lon2 - lon1)
	var a = math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(deg2rad(lat1))*math.Cos(deg2rad(lat2))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var d = R * c // Distance in km
	return d
}

func deg2rad(deg float32) float64 {
	return float64(deg * (math.Pi / 180))
}
