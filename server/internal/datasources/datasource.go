package datasources

import "context"

type BlahajAvailability struct {
	StoreID      string         `json:"store_id"`
	StoreName    string         `json:"store_name"`
	StoreCountry string         `json:"store_country"`
	Location     GeoCoordinates `json:"location"`
	Number       int64          `json:"number"`
}

type GeoCoordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type CountryAvailability struct {
	Country string `json:"country"`
	Number  int64  `json:"number"`
}

type Datasource interface {
	GlobalAvailability(ctx context.Context) ([]CountryAvailability, error)
	Availability(ctx context.Context) ([]BlahajAvailability, error)
}
