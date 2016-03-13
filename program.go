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
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	timeout := time.Duration(500 * time.Millisecond)
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	host := os.Getenv("HOST")
	println("CouchDB Address ", host, port)
	conn, err := couchdb.NewConnection(host, port, timeout)
	if err != nil {
		panic(err.Error())
	}

	db := conn.SelectDB("mydb", nil)

	person.CreatePersonService(router, db)
	health.CreateHealthService(router, conn)
	router.Run("0.0.0.0:8080")
}
