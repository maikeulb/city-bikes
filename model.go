package main

type NetworksResponse struct {
	Networks []Network `json:"networks"`
}

type NetworkResponse struct {
	Network Network `json:"network"`
}

type Network struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

type Location struct {
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
