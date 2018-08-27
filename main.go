package main

import (
  "net/http"
  "log"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/kevinmoran100/arqui2_practica1/Cassandra"
  "github.com/kevinmoran100/arqui2_practica1/Data"
)

type heartbeatResponse struct {
  Status string `json:"status"`
  Code int `json:"code"`
}

func main() {
  CassandraSession := Cassandra.Session
  defer CassandraSession.Close()

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", heartbeat)
  router.HandleFunc("/data/new",Data.Post)
  router.HandleFunc("/data", Data.Get)
  router.HandleFunc("/data/{fecha}", Data.GetOne)
  log.Fatal(http.ListenAndServe(":8080", router))
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
