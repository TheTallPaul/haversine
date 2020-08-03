/*
Package haversine calculates the distance between to Latitude/Longitude pairs,
taking into account the curvature of the Earth. See
https://en.wikipedia.org/wiki/Haversine_formula for more information on the
formula.

haversine supports popular Lat/Long structs, so converting your structs for the
function is unnecessary.

Simply pass two supported types to HaversineDistance() to get the distance in
meters:

  package main
  import (
    "fmt"

    HaversineDistance "github.com/TheTallPaul/haversine"
    "google.golang.org/genproto/googleapis/type/latlng"
  )
  func main() {
    zeroPoint := latlng.LatLng{Latitude: 0.0, Longitude: 0.0}
    onePoint := latlng.LatLng{Latitude: 1.0, Longitude: -1.0}

    distance := HaversineDistance(&zeroPoint, &onePoint)
    fmt.Printf("Distance between (%v, %v) and (%v, %v): %v meters",
      zeroPoint.Latitude,
      zeroPoint.Longitude,
      onePoint.Latitude,
      onePoint.Longitude,
      distance
    );
  }
Output:
  Distance between (0, 0) and (1, -1): 157249.38127194397 meters
*/

package haversine

import (
	"math"

	"google.golang.org/genproto/googleapis/type/latlng"
	"googlemaps.github.io/maps"
)

// EarthRadius is the volumetric mean radius of the Earth in meters taken from
// https://nssdc.gsfc.nasa.gov/planetary/factsheet/earthfact.html
const EarthRadius = 6371000.0
const degreesToRadians = math.Pi / 180

/*
Distance finds the Haversine distance in meters between two points. It
accepts the following types:

* google.golang.org/genproto/googleapis/type/latlng (*latlng.LatLng)

* googlemaps.github.io/maps (*maps.LatLng)
*/
func HaversineDistance(pointA, pointB interface{}) float64 {
	// *latlng.LatLng
	{
		newPointA, aIsType := pointA.(*latlng.LatLng)
		newPointB, bIsType := pointB.(*latlng.LatLng)
		if aIsType && bIsType {
			return HaversineFormula(
				newPointA.Latitude,
				newPointA.Longitude,
				newPointB.Latitude,
				newPointB.Longitude)
		}
	}

	// *maps.LatLng
	{
		newPointA, aIsType := pointA.(*maps.LatLng)
		newPointB, bIsType := pointB.(*maps.LatLng)
		if aIsType && bIsType {
			return HaversineFormula(newPointA.Lat, newPointA.Lng,
				newPointB.Lat, newPointB.Lng)
		}
	}

	return 0
}

// HaversineFormula finds the Haversine distance from the latitude and longitude
// pairs of A and B. Use this if you don't want to use any structs to find the
// distance between two points.
func HaversineFormula(latA, longA, latB, longB float64) float64 {
	latA *= degreesToRadians
	latB *= degreesToRadians
	latDist := latB - latA
	longDist := (longB - longA) * degreesToRadians

	return 2 * EarthRadius * math.Asin(math.Sqrt(math.Pow(
		math.Sin(latDist/2), 2)+math.Pow(math.Sin(longDist/2), 2)*
		math.Cos(latA)*math.Cos(latB)))
}
