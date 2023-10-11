package gis

import (
	"encoding/json"
	"github.com/violin8/pkg/cast"
	"math"
)

type Location struct {
	Longitude string `json:"longitude"` // 经度
	Latitude  string `json:"latitude"`  // 纬度
}

func NewLocation(locationStr string) Location {
	l := Location{}
	_ = json.Unmarshal([]byte(locationStr), &l)
	return l
}

func (loc Location) Lng() float64 {
	return cast.ToFloat64(loc.Longitude)
}

func (loc Location) Lat() float64 {
	return cast.ToFloat64(loc.Latitude)
}

// Sphere 两经纬度坐标距离
// lonA, latA分别为A点的纬度和经度
// lonB, latB分别为B点的纬度和经度
// 单位: 米
func Sphere(lat1, lng1, lat2, lng2 float64) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := PI * lat1 / 180
	radlat2 := PI * lat2 / 180

	theta := lng1 - lng2
	radtheta := PI * theta / 180

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	return dist * 1.609344 * 1000
}
