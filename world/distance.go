package world

import "math"

const (
	earthRadius  = 6372.8   // Earth's radius in kilometers
	speedOfLight = 200000.0 // Speed of light in fiber-optic cables in kilometers per second
)

// Coordinate holds the Latitude and Longitude of a point on Earth.
type Coordinate struct {
	Latitude  float64
	Longitude float64
}

// ToRadians converts degrees to radians.
func ToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

// HaversineDistance calculates the distance in KM between two coordinates on Earth.
func HaversineDistance(coord1, coord2 Coordinate) float64 {
	lat1Rad := ToRadians(coord1.Latitude)
	lng1Rad := ToRadians(coord1.Longitude)
	lat2Rad := ToRadians(coord2.Latitude)
	lng2Rad := ToRadians(coord2.Longitude)

	deltaLat := lat2Rad - lat1Rad
	deltaLng := lng2Rad - lng1Rad

	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(deltaLng/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

// MinimumLatency calculates the minimum communications latency in seconds.
// Expects a distance given in kilometers.
func MinimumLatency(distance float64) float64 {
	return distance / speedOfLight
}
