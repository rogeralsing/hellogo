package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rhinoman/couchdb-go"
	"github.com/rogeralsing/hellogo/util"
	"time"
	"net/http"
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
	var timeout = time.Duration(500 * time.Millisecond)
	conn, _ := couchdb.NewConnection("127.0.0.1", 5984, timeout)
	//auth := couchdb.BasicAuth{Username: "user", Password: "password" }
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

		if err := c.BindJSON(&doc);err == nil {
			db.Save(doc, id, rev)
			c.JSON(http.StatusOK, doc)
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	})

	router.Run("0.0.0.0:8080")
}
