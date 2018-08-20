package Data
//
// import (
// 	"github.com/gocql/gocql"
// )

// Data struct to hold profile data for our Data
type Data struct {
	Fecha        string `json:"fecha"`
	Humedad      string `json:"humedad"`
	Coordenadas  string `json:"coordenadas"`
	Presion      string `json:"presion"`
	Radiacion    string `json:"radiacion"`
	Temperatura  string `json:"temperatura"`
	Viento       string `json:"viento"`
}

// GetDataResponse to form payload returning a single Data struct
type GetDataResponse struct {
	Dato Data `json:"dato"`
}

// AllDatasResponse to form payload of an array of Data structs
type AllDataResponse struct {
	Datos []Data `json:"datos"`
}

// NewDataResponse builds a payload of new Data resource ID
type NewDataResponse struct {
	Fecha string `json:"fecha"`
}

// ErrorResponse returns an array of error strings if appropriate
type ErrorResponse struct {
	Errors []string `json:"errors"`
}
