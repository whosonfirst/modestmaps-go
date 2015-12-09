package modestmaps

import (
	"fmt"
)

func NewLocation(lat float64, lon float64) *Location {
	return &Location{lat, lon}
}

type Location struct {
	lat float64
	lon float64
}

func (l *Location) Lat() float64 {
	return l.lat
}

func (l *Location) Lon() float64 {
	return l.lon
}

func (l *Location) Stringer() string {
	return fmt.Sprintf("%.3f, %.3f", l.lat, l.lon)
}
