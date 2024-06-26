package db

type QandA struct {
	Question    string
	Description string
}

type SafariBreakDown struct {
	Day      int16
	Title    string
	Activity string
	Image    string
}

type Safari struct {
	ID          int32
	Title       string
	Guide       string
	Days        int16
	Description string
	Image       string
	Breakdown   []SafariBreakDown
}

type Vehicle struct {
	ID         int16
	Name       string
	Passengers int16
	Suitcases  int16
	Type       string
	PricePerKM string
	Images     []string
	Make       VehicleMake
}

type VehicleMake struct {
	Engine        string
	InteriorColor string
	Power         string
	FuelType      string
	Length        string
	ExteriorColor string
	Transmission  string
	Extras        string
}