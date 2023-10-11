package gis

import "testing"

func TestGetDistance(t *testing.T) {
	lat1 := 40.003975
	lng1 := 116.539949

	lat2 := 31.777769
	lng2 := 117.404472
	t.Logf(`1: %v`, Sphere(lat1, lng1, lat2, lng2))
}
