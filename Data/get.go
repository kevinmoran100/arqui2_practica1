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
  var dataList []Data
  m := map[string]interface{}{}

  query := "SELECT * FROM data"
  iterable := Cassandra.Session.Query(query).Iter()
  for iterable.MapScan(m) {
    dataList = append(dataList, Data{
      fecha: m["fecha"].(time.Time).String(),
      humedad: m["humedad"].(string),
      coordenadas: m["coordenadas"].(string),
      presion: m["presion"].(string),
      radiacion: m["radiacion"].(string),
      temperatura: m["temperatura"].(string),
      viento: m["viento"].(string),
    })
    m = map[string]interface{}{}
  }

  json.NewEncoder(w).Encode(AllDataResponse{Datos: dataList})
}

func GetOne(w http.ResponseWriter, r *http.Request) {
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
        fecha: m["fecha"].(time.Time).String(),
        humedad: m["humedad"].(string),
        coordenadas: m["coordenadas"].(string),
        radiacion: m["radiacion"].(string),
        temperatura: m["temperatura"].(string),
        presion: m["presion"].(string),
        viento: m["viento"].(string),
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
