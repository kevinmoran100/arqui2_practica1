package Cassandra

import (
  "github.com/gocql/gocql"
  "fmt"
)
var Session *gocql.Session

func init() {
  var err error

  cluster := gocql.NewCluster("35.192.151.177")
  cluster.Keyspace = "practica1"
  Session, err = cluster.CreateSession()
  if err != nil {
    panic(err)
  }
  fmt.Println("cassandra init done")
}
