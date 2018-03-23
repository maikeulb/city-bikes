package main

type Response struct {
	Networks []Network `json:"networks"`
}

type Network struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

type Location struct {
	City      int     `json:"city"`
	Country   int     `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
