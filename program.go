package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rhinoman/couchdb-go"
	"github.com/rogeralsing/hellogo/util"
	"net/http"
	"os"
	"strconv"
	"time"
)

type PersonDocument struct {
	Name     string   `form:"name" json:"name"`
	Age      int      `form:"age" json:"age"`
	Children []string `"form:children" json:"children"`
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	util.Hello()
	timeout := time.Duration(500 * time.Millisecond)
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	host := os.Getenv("HOST")
	println("CouchDB Address ", host, port)

	if conn, err := couchdb.NewConnection(host, port, timeout); err == nil {
		db := conn.SelectDB("mydb", nil)

		router.GET("/bar/:id", func(c *gin.Context) {
			id := c.Param("id")
			var doc PersonDocument
			db.Read(id, &doc, nil)
			c.JSON(http.StatusOK, doc)
		})

		router.PUT("/bar/:id", func(c *gin.Context) {
			id := c.Param("id")
			var doc PersonDocument
			rev, _ := db.Read(id, &doc, nil)

			if err := c.BindJSON(&doc); err == nil {
				db.Save(doc, id, rev)
				c.JSON(http.StatusOK, doc)
			} else {
				c.JSON(http.StatusBadRequest, err.Error())
			}
		})

		router.Run("0.0.0.0:8080")
	} else {
		panic(err.Error())
	}
	//auth := couchdb.BasicAuth{Username: "user", Password: "password" }

}
