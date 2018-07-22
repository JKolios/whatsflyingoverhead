package dump1090fa

import (
	"math"

	"github.com/paulmach/go.geo"
)

const FEET_TO_METERS = 0.3048

type AircraftFile struct {
	Now      float64    `json:"now"`
	Messages int        `json:"messages"`
	Aircraft []Aircraft `json:"aircraft"`
}

type Aircraft struct {
	Hex      string        `json:"hex"`
	Squawk   string        `json:"squawk,omitempty"`
	Flight   string        `json:"flight,omitempty"`
	Lat      float64       `json:"lat,omitempty"`
	Lon      float64       `json:"lon,omitempty"`
	Nucp     int           `json:"nucp,omitempty"`
	SeenPos  float64       `json:"seen_pos,omitempty"`
	Altitude int           `json:"altitude,omitempty"`
	VertRate int           `json:"vert_rate,omitempty"`
	Track    int           `json:"track,omitempty"`
	Speed    int           `json:"speed,omitempty"`
	Category string        `json:"category,omitempty"`
	Mlat     []interface{} `json:"mlat"`
	Tisb     []interface{} `json:"tisb"`
	Messages int           `json:"messages"`
	Seen     float64       `json:"seen"`
	Rssi     float64       `json:"rssi"`
}

// Distance calculates the three dimensional distance between an aircraft and a given geographical point.
// The difference of heights between the point and the aircraft is taken into account
func (aircraft Aircraft) Distance(fromLatitude, fromLongitude, fromAltitude float64) float64 {
	fromPoint := geo.NewPointFromLatLng(fromLatitude, fromLongitude)
	aircraftPoint := geo.NewPointFromLatLng(aircraft.Lat, aircraft.Lon)
	haversineDistance := fromPoint.GeoDistanceFrom(aircraftPoint, true)
	altitudeDifference := aircraft.altitudeInMeters() - fromAltitude
	return linearDistance(haversineDistance, altitudeDifference)
}

// HasCoordinates checks if we know enough about an aircraft's location to determine its distance from the receiver
func (aircraft Aircraft) HasCoordinates() bool {
	return aircraft.Lat != 0 && aircraft.Lon != 0 && aircraft.Altitude != 0
}

func (aircraft Aircraft) altitudeInMeters() float64 {
	return feetToMeters(float64(aircraft.Altitude))
}

func linearDistance(haversineDistance, altitudeDifference float64) float64 {
	return math.Sqrt(math.Pow(haversineDistance, 2) + math.Pow(altitudeDifference, 2))
}

func feetToMeters(feet float64) float64 {
	return feet * FEET_TO_METERS
}
