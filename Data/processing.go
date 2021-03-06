package Data

import (
	"net/http"
	// "fmt"
	// "strconv"
)

// FormToData -- fills a Data struct with submitted form data
// params:
// r - request reader to fetch form data or url params (unused here)
// returns:
// Data struct if successful
// array of strings of errors if any occur during processing
func FormToData(r *http.Request) (Data, []string) {
	var Data Data
	var errStr string
	var errs []string
	// var err error
	// fmt.Println(r)
	Data.Fecha, errStr = processFormField(r, "fecha")
	errs = appendError(errs, errStr)
	Data.Humedad, errStr = processFormField(r, "humedad")
	errs = appendError(errs, errStr)
	Data.Coordenadas, errStr = processFormField(r, "coordenadas")
	errs = appendError(errs, errStr)
  Data.Presion, errStr = processFormField(r, "presion")
  errs = appendError(errs, errStr)
	Data.Radiacion, errStr = processFormField(r, "radiacion")
	errs = appendError(errs, errStr)
	Data.Temperatura, errStr = processFormField(r, "temperatura")
	errs = appendError(errs, errStr)
	Data.Viento, errStr = processFormField(r, "viento")
	errs = appendError(errs, errStr)


	// ageStr, errStr = processFormField(r, "age")
	// if len(errStr) != 0 {
	// 	errs = append(errs, errStr)
	// } else {
	// 	Data.Age, err = strconv.Atoi(ageStr)
	// 	if err != nil {
	// 		errs = append(errs, "Parameter 'age' not an integer")
	// 	}
	// }
	return Data, errs
}

func appendError(errs []string, errStr string) ([]string) {
	if len(errStr) > 0 {
		errs = append(errs, errStr)
	}
	return errs
}

func processFormField(r *http.Request, field string) (string, string) {
	fieldData := r.PostFormValue(field)
	if len(fieldData) == 0 {
		return "", "Missing '" + field + "' parameter, cannot continue"
	}
	return fieldData, ""
}
