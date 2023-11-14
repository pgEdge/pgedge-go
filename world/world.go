package world

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed regions.json
var regionsData string

//go:embed locations.json
var locationsData string

var regions []*Region

var locations []*Location

func init() {
	type region struct {
		Cloud    string   `json:"cloud"`
		Name     string   `json:"name"`
		Zones    []string `json:"zones"`
		Location string   `json:"location"`
	}
	var rawRegions []*region
	if err := json.Unmarshal([]byte(regionsData), &rawRegions); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(locationsData), &locations); err != nil {
		panic(err)
	}
	for _, r := range rawRegions {
		loc, ok := GetLocation(r.Location)
		if !ok {
			panic("unknown location: " + r.Location)
		}
		var zoneNames []string
		for _, z := range r.Zones {
			zoneNames = append(zoneNames, fmt.Sprintf("%s%s", r.Name, z))
		}
		regions = append(regions, &Region{
			Cloud:    r.Cloud,
			Name:     r.Name,
			Zones:    zoneNames,
			Location: loc,
		})
	}
}

type Location struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Name       string  `json:"name,omitempty"`        // "Northern Virginia"
	Code       string  `json:"code,omitempty"`        // "IAD"
	Country    string  `json:"country,omitempty"`     // "US"
	City       string  `json:"city,omitempty"`        // "Sterling"
	MetroCode  string  `json:"metro_code,omitempty"`  // "511"
	PostalCode string  `json:"postal_code,omitempty"` // "20165"
	Region     string  `json:"region,omitempty"`      // "Virginia"
	RegionCode string  `json:"region_code,omitempty"` // "VA"
	Timezone   string  `json:"timezone,omitempty"`    // "America/New_York"
}

func (l *Location) Coordinate() Coordinate {
	return Coordinate{Latitude: l.Latitude, Longitude: l.Longitude}
}

type Region struct {
	Cloud    string    `json:"cloud"` // "aws"
	Name     string    `json:"name"`  // "us-east-2"
	Zones    []string  `json:"zones"` // ["us-east-2a", "us-east-2b", ...]
	Location *Location `json:"location"`
}

func Regions() []*Region {
	return regions
}

func Locations() []*Location {
	return locations
}

func GetRegion(cloud, name string) (*Region, bool) {
	for _, r := range regions {
		if r.Cloud == cloud && r.Name == name {
			return r, true
		}
	}
	return nil, false
}

func GetLocation(code string) (*Location, bool) {
	for _, l := range locations {
		if l.Code == code {
			return l, true
		}
	}
	return nil, false
}
