package Data

import (
"net/http"
"github.com/gocql/gocql"
"encoding/json"
"github.com/kevinmoran100/arqui2_practica1/Cassandra"
"fmt"
)

func Post(w http.ResponseWriter, r *http.Request) {
  var errs []string
  // var fecha string

  // FormToData() is included in Datas/processing.go
  // we will describe this later
  Data, errs := FormToData(r)

  // have we created a Data correctly
  var created bool = false

  // if we had no errors from FormToData, we will
  // attempt to save our data to Cassandra
  if len(errs) == 0 {
    fmt.Println("creating a new Data")

    // generate a unique UUID for this Data
    // gocqlUuid = gocql.TimeUUID()

    // write data to Cassandra
    if err := Cassandra.Session.Query(`
      INSERT INTO Data (fecha,humedad,coordenadas,radiacion,temperatura,presion,viento) VALUES (?, ?, ?, ?, ?, ?)`,
      Data.fecha, Data.humedad, Data.coordenadas, Data.radiacion, Data.temperatura, Data.presion, Data.viento).Exec(); err != nil {
      errs = append(errs, err.Error())
    } else {
      created = true
    }
  }

  // depending on whether we created the Data, return the
  // resource ID in a JSON payload, or return our errors
  if created {
    fmt.Println("fecha", Data.fecha)
    json.NewEncoder(w).Encode(NewDataResponse{fecha: Data.fecha})
  } else {
    fmt.Println("errors", errs)
    json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
  }
}
