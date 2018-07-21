package dump1090fa

import "time"

type AircraftFile struct {
	Now      time.Time  `json:"now"`
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
