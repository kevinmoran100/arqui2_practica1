package Data

import (
  "net/http"
  "github.com/gocql/gocql"
  "encoding/json"
  "github.com/kevinmoran100/arqui2_practica1/Cassandra"
  "github.com/gorilla/mux"
  "time"
  "strconv"
  "fmt"
)

func Get(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  var dataList []Data
  m := map[string]interface{}{}

  query := "SELECT * FROM data"
  iterable := Cassandra.Session.Query(query).Iter()
  for iterable.MapScan(m) {
    dataList = append(dataList, Data{
      Fecha: m["fecha"].(time.Time).String(),
      Humedad: m["humedad"].(string),
      Coordenadas: m["coordenadas"].(string),
      Presion: m["presion"].(string),
      Radiacion: m["radiacion"].(string),
      Temperatura: m["temperatura"].(string),
      Viento: m["viento"].(string),
    })
    m = map[string]interface{}{}
  }

  json.NewEncoder(w).Encode(AllDataResponse{Datos: dataList})
}

func GetOne(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  var data Data
  var errs []string
  var found bool = false

  vars := mux.Vars(r)
  id := vars["fecha"]
  i, err := strconv.ParseInt(id, 10, 64)
  if err != nil {
      panic(err)
  }
  tm := time.Unix(i, 0)
    m := map[string]interface{}{}
    query := "SELECT * FROM data WHERE fecha=? LIMIT 1"
    iterable := Cassandra.Session.Query(query, tm).Consistency(gocql.One).Iter()
    for iterable.MapScan(m) {
      found = true
      data = Data{
        Fecha: m["fecha"].(time.Time).String(),
        Humedad: m["humedad"].(string),
        Coordenadas: m["coordenadas"].(string),
        Presion: m["presion"].(string),
        Radiacion: m["radiacion"].(string),
        Temperatura: m["temperatura"].(string),
        Viento: m["viento"].(string),
      }
    }
    if !found {
      errs = append(errs, "data not found")
    }


  if found {
    fmt.Println(data);
    json.NewEncoder(w).Encode(GetDataResponse{Dato: data})
  } else {
    json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
  }
}
