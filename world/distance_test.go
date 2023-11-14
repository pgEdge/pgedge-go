package world

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHaversineDistance(t *testing.T) {
	// haversine(36.12, -86.67, 33.94, -118.40)
	// 2887.26
	// https://rosettacode.org/wiki/Haversine_formula
	coord1 := Coordinate{36.12, -86.67}
	coord2 := Coordinate{33.94, -118.40}
	expected := 2887
	actual := int(HaversineDistance(coord1, coord2))
	require.Equal(t, expected, actual, "HaversineDistance()")
}

func TestHaversineDistance2(t *testing.T) {
	// NEBRASKA, USA (Latitude : 41.507483, longitude : -99.436554)
	// KANSAS, USA (Latitude : 38.504048, Longitude : -98.315949)
	coord1 := Coordinate{41.507483, -99.436554}
	coord2 := Coordinate{38.504048, -98.315949}
	expected := 347
	actual := int(HaversineDistance(coord1, coord2))
	require.Equal(t, expected, actual, "HaversineDistance()")
}

func TestDublinToLondon(t *testing.T) {
	dub := Coordinate{53.426448, -6.24991}
	lon := Coordinate{51.5072, -0.1276}
	expected := 466
	actual := int(HaversineDistance(dub, lon))
	require.Equal(t, expected, actual, "HaversineDistance()")
}
