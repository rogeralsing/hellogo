package main

import (
	"github.com/rogeralsing/hellogo/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/rogeralsing/hellogo/util"
	"net/http"
	"os"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	util.Hello()

	router.GET("/foo", func(c *gin.Context) {

		res,_ := os.Hostname()
		c.String(http.StatusOK, "Hello %s", res)
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
