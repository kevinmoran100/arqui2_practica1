package Data

import (
	"github.com/gocql/gocql"
)

// Data struct to hold profile data for our Data
type Data struct {
	fecha        string `json:"fecha"`
	humedad      string `json:"humedad"`
	coordenadas  string `json:"coordenadas"`
	presion      string `json:"presion"`
	radiacion    string `json:"radiacion"`
	temperatura  string `json:"temperatura"`
	viento       string `json:"viento"`
}

// GetDataResponse to form payload returning a single Data struct
type GetDataResponse struct {
	Data Data `json:"Data"`
}

// AllDatasResponse to form payload of an array of Data structs
type AllDatasResponse struct {
	Datas []Data `json:"Datas"`
}

// NewDataResponse builds a payload of new Data resource ID
type NewDataResponse struct {
	fecha string`json:"fecha"`
}

// ErrorResponse returns an array of error strings if appropriate
type ErrorResponse struct {
	Errors []string `json:"errors"`
}
