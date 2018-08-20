import (
  "net/http"
  "github.com/gocql/gocql"
  "encoding/json"
  "github.com/dataname/projectname/Cassandra"
  "github.com/gorilla/mux"
  "fmt"
)

func Get(w http.ResponseWriter, r *http.Request) {
  var dataList []data
  m := map[string]interface{}{}

  query := "SELECT * FROM data"
  iterable := Cassandra.Session.Query(query).Iter()
  for iterable.MapScan(m) {
    dataList = append(dataList, data{
      fecha: m["fecha"].(string),
      humedad: m["humedad"].(string),
      coordenadas: m["coordenadas"].(string),
      radiacion: m["radiacion"].(string),
      temperatura: m["temperatura"].(string),
      presion: m["presion"].(string),
      viento: m["viento"].(string),
    })
    m = map[string]interface{}{}
  }

  json.NewEncoder(w).Encode(datasResponse{datas: dataList})
}

func GetOne(w http.ResponseWriter, r *http.Request) {
  var data data
  var errs []string
  var found bool = false

  vars := mux.Vars(r)
  id := vars["fecha"]


    m := map[string]interface{}{}
    query := "SELECT * FROM data WHERE fecha=? LIMIT 1"
    iterable := Cassandra.Session.Query(query, uuid).Consistency(gocql.One).Iter()
    for iterable.MapScan(m) {
      found = true
      data = data{
        fecha: m["fecha"].(string),
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
    json.NewEncoder(w).Encode(GetdataResponse{data: data})
  } else {
    json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
  }
}
