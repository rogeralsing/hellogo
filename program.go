package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rhinoman/couchdb-go"
	"github.com/rogeralsing/hellogo/util"
	"net/http"
	"os"
	"time"
)

type PersonDocument struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Children []string `json:"children"`
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
		var res PersonDocument
		db.Read(id, &res, nil)
		c.JSON(200, res)
	})

	router.GET("/foo", func(c *gin.Context) {

		theDoc := PersonDocument{
			Name:     "Olle",
			Age:      80,
			Children: []string{"Arne", "Pelle", "Eva"},
		}

		theId := "mittid123"
		//The third argument here would be a revision, if you were updating an existing document
		_, _ = db.Save(theDoc, theId, "")

		var x PersonDocument
		db.Read(theId, &x, nil)

		res, _ := os.Hostname()
		c.String(http.StatusOK, "Hello %s %s", res, x.Children[0])
	})

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run("0.0.0.0:8080")
}
