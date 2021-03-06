package Data

import (
"net/http"
"encoding/json"
"github.com/kevinmoran100/arqui2_practica1/Cassandra"
"fmt"
"time"
"strconv"
)

func Post(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  var errs []string
  // var fecha string

  // FormToData() is included in Datas/processing.go
  // we will describe this later
  Data, errs := FormToData(r)

  // have we created a Data correctly
  var created bool = false

  // if we had no errors from FormToData, we will
  // attempt to save our data to Cassandra
  //convertir fecha
  i, err := strconv.ParseInt(Data.Fecha, 10, 64)
  if err != nil {
      panic(err)
  }
  tm := time.Unix(i, 0)
  // fmt.Println(tm)

  if len(errs) == 0 {
    fmt.Println("creating a new Data")

    // generate a unique UUID for this Data
    // gocqlUuid = gocql.TimeUUID()

    // write data to Cassandra
    // fmt.Println(Data)
    if err := Cassandra.Session.Query(`
      INSERT INTO practica1.Data (fecha,humedad,coordenadas,radiacion,temperatura,presion,viento) VALUES (?, ?, ?, ?, ?, ?,?)`,
      tm, Data.Humedad, Data.Coordenadas, Data.Radiacion, Data.Temperatura, Data.Presion, Data.Viento).Exec(); err != nil {
      errs = append(errs, err.Error())
    } else {
      created = true
    }
  }

  // depending on whether we created the Data, return the
  // resource ID in a JSON payload, or return our errors
  if created {
    // fmt.Println("fecha", Data.fecha)
    // json.NewEncoder(w).Encode("{\"fecha\":\""+Data.fecha+"\"}")
    json.NewEncoder(w).Encode(NewDataResponse{Fecha: Data.Fecha})
  } else {
    // fmt.Println("errors", errs)
    json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
  }
}
