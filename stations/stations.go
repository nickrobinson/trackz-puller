package stations

import "fmt"

type StationInfo struct {
	Ssid   string `json:"ssid"`
	Bssid  string `json:"bssid"`
	Signal int32  `json:"signal"`
}

type Stations struct {
	Stations []StationInfo `json:"stations"`
}

func init() {
	fmt.Println("Initializing Stations")
}
