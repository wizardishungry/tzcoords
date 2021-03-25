package tzcoords

import (
	"fmt"
	"time"
)

// ErrNotFound indicates that coordinates for a timezone are unavailable
var ErrNotFound = fmt.Errorf("not found")

// ByLocation retreives a lat/lot for a Location
func ByLocation(loc *time.Location) (float64, float64, error) {
	return ByString(loc.String())
}

// ByLocation retreives a lat/lot for a IANA time zone identifier
func ByString(s string) (float64, float64, error) {
	ll, ok := toCoords[s]
	if !ok {
		return 0, 0, fmt.Errorf("%s %w", s, ErrNotFound)
	}
	return ll.Lat, ll.Lon, nil
}

//go:generate go run ./gen/gen.go -output coords.go
