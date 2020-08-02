package haversine

import (
	"math"

	"google.golang.org/genproto/googleapis/type/latlng"
)

// EarthRadius is the volumetric mean radius of the Earth in meters taken from
// https://nssdc.gsfc.nasa.gov/planetary/factsheet/earthfact.html
const EarthRadius = 6371000.0
const degreesToRadians = math.Pi / 180



// Haversine finds the Haversine distance in meters between two points
func Haversine(pointA, pointB *latlng.LatLng) float64 {
	latDist := (pointB.Latitude - pointA.Latitude) * degreesToRadians
	lngDist := (pointB.Longitude - pointA.Longitude) * degreesToRadians

	latA := pointA.Latitude * degreesToRadians
	latB := pointB.Latitude * degreesToRadians

	return 2 * EarthRadius * math.Asin(math.Sqrt(
		math.Pow(math.Sin(latDist/2), 2)+
			math.Pow(math.Sin(lngDist/2), 2)*
			math.Cos(latA)*math.Cos(latB)))
}
