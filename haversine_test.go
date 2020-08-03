package haversine

import (
	"testing"

	"github.com/adrianmo/go-nmea"
	"google.golang.org/genproto/googleapis/type/latlng"
	"googlemaps.github.io/maps"
)

// googleapis
var zeroPoint = latlng.LatLng{Latitude: 0.0, Longitude: 0.0}
var onePoint = latlng.LatLng{Latitude: 1.0, Longitude: -1.0}
var broadwayYamhill = latlng.LatLng{Latitude: 45.518673, Longitude: -122.679996}
var pioneerCourthouse = latlng.LatLng{
	Latitude:  45.51867299282918,
	Longitude: -122.67871393198214,
}

var haversineDistanceTestCases = []struct {
	pointA           *latlng.LatLng
	pointB           *latlng.LatLng
	expectedDistance float64
}{
	{
		&zeroPoint,
		&onePoint,
		157249.38127194397,
	},
	{
		&broadwayYamhill,
		&pioneerCourthouse,
		99.88810211965017,
	},
}

// maps
var zeroPointMaps = maps.LatLng{Lat: 0.0, Lng: 0.0}
var onePointMaps = maps.LatLng{Lat: 1.0, Lng: -1.0}
var broadwayYamhillMaps = maps.LatLng{Lat: 45.518673, Lng: -122.679996}
var pioneerCourthouseMaps = maps.LatLng{
	Lat: 45.51867299282918,
	Lng: -122.67871393198214,
}

var haversineDistanceTestCasesMaps = []struct {
	pointA           *maps.LatLng
	pointB           *maps.LatLng
	expectedDistance float64
}{
	{
		&zeroPointMaps,
		&onePointMaps,
		157249.38127194397,
	},
	{
		&broadwayYamhillMaps,
		&pioneerCourthouseMaps,
		99.88810211965017,
	},
}

// nmea
var sentence = nmea.BaseSentence{}
var zeroPointGLL = nmea.GLL{
	BaseSentence: sentence,
	Latitude:     0.0,
	Longitude:    0.0,
}
var onePointGLL = nmea.GLL{
	BaseSentence: sentence,
	Latitude:     1.0,
	Longitude:    1.0,
}

var haversineDistanceTestCasesNMEA = []struct {
	pointA           nmea.GLL
	pointB           nmea.GLL
	expectedDistance float64
}{
	{
		zeroPointGLL,
		onePointGLL,
		157249.38127194397,
	},
}

func TestHaversineDistance(t *testing.T) {
	for _, input := range haversineDistanceTestCases {
		distance := HaversineDistance(input.pointA, input.pointB)

		if distance != input.expectedDistance {
			t.Errorf(
				"FAIL: Want distance from (%v, %v) to (%v, %v"+
					") to be: %v m but we got %v m ("+
					"latlng package)",
				input.pointA.Latitude,
				input.pointA.Longitude,
				input.pointB.Latitude,
				input.pointB.Longitude,
				input.expectedDistance,
				distance,
			)
		}
	}

	for _, input := range haversineDistanceTestCasesMaps {
		distance := HaversineDistance(input.pointA, input.pointB)

		if distance != input.expectedDistance {
			t.Errorf(
				"FAIL: Want distance from (%v, %v) to (%v, %v"+
					") to be: %v m but we got %v m (maps "+
					"package)",
				input.pointA.Lat,
				input.pointA.Lng,
				input.pointB.Lat,
				input.pointB.Lng,
				input.expectedDistance,
				distance,
			)
		}
	}

	for _, input := range haversineDistanceTestCasesNMEA {
		distance := HaversineDistance(input.pointA, input.pointB)

		if distance != input.expectedDistance {
			t.Errorf(
				"FAIL: Want distance from (%v, %v) to (%v, %v"+
					") to be: %v m but we got %v m (nmea "+
					"package)",
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
