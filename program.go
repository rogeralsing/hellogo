package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rhinoman/couchdb-go"
	"net/http"
	"os"
	"strconv"
	"time"
)

type PersonDocument struct {
	Name     string   `form:"name" json:"name"`
	Age      int      `form:"age" json:"age"`
	Children []string `form:"children" json:"children"`
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	timeout := time.Duration(500 * time.Millisecond)
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	host := os.Getenv("HOST")
	println("CouchDB Address ", host, port)
	conn, err := couchdb.NewConnection(host, port, timeout);
	if err != nil {
		panic(err.Error())
	}

	db := conn.SelectDB("mydb", nil)

	router.GET("/bar/:id", func(c *gin.Context) {
		id := c.Param("id")
		var doc PersonDocument
		if _, err := db.Read(id, &doc, nil); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, doc)
	})

	router.PUT("/bar/:id", func(c *gin.Context) {
		id := c.Param("id")
		var doc PersonDocument
		rev, _ := db.Read(id, &doc, nil)

		if err := c.BindJSON(&doc); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if _, err := db.Save(doc, id, rev); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, doc)
	})

	router.Run("0.0.0.0:8080")
}
