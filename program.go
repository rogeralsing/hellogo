package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rhinoman/couchdb-go"
	"github.com/rogeralsing/hellogo/health"
	"github.com/rogeralsing/hellogo/person"
	"os"
	"strconv"
	"time"
)

func main() {
	dbTimeout := time.Duration(500 * time.Millisecond)
	dbPort, _ := strconv.Atoi(os.Getenv("PORT"))
	if dbPort == 0 {
		panic("No DB Port given PORT")
	}
	dbHost := os.Getenv("HOST")
	if dbHost == "" {
		panic("No DB Host given $HOST")
	}
	println("CouchDB Address ", dbHost, dbPort)
	conn, err := couchdb.NewConnection(dbHost, dbPort, dbTimeout)
	if err != nil {
		panic(err.Error())
	}

	db := conn.SelectDB("mydb", nil)

	router := gin.Default()
	personDB := person.CouchDBPersonRepository{DB: db}
	person.CreatePersonService(router, personDB)
	health.CreateHealthService(router, conn)
	router.Run("0.0.0.0:8080")
}
