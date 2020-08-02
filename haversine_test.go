package haversine

import (
  "testing"
  "google.golang.org/genproto/googleapis/type/latlng"
)

var zeroPoint = latlng.LatLng{Latitude: 0.0, Longitude: 0.0}
var onePoint = latlng.LatLng{Latitude: 1.0, Longitude: -1.0}
var broadwayYamhill = latlng.LatLng{Latitude: 45.518673, Longitude: -122.679996}
var pioneerCourthouse = latlng.LatLng{
	Latitude:  45.51867299282918,
	Longitude: -122.67871393198214,
}

var haversineTestCases = []struct {
	pointA           *latlng.LatLng
	pointB           *latlng.LatLng
	expectedDistance float64
}{
	{
		&zeroPoint,
		&onePoint,
		157425.537108412,
	},
	{
		&broadwayYamhill,
		&pioneerCourthouse,
		99.99999999829213,
	},
}

func TestHaversine(t *testing.T) {
	for _, input := range haversineTestCases {
		distance := haversine(input.pointA, input.pointB)

		if distance != input.expectedDistance {
			t.Errorf(
				"FAIL: Want distance from (%v, %v) to (%v, %v) to be: "+
					"%v but we got %v",
				input.pointA.Latitude,
				input.pointA.Longitude,
				input.pointB.Latitude,
				input.pointB.Longitude,
				input.expectedDistance,
				distance,
			)
		}
	}
}
