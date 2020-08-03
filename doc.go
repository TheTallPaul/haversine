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
