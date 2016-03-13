package health

import (
	"github.com/gin-gonic/gin"
	"github.com/rhinoman/couchdb-go"
	"net/http"
)

func CreateHealthService(router *gin.Engine, conn *couchdb.Connection) {
	health := router.Group("/health")
	{
		health.GET("", func(c *gin.Context) {
			//check DB
			if err := conn.Ping(); err == nil {
				c.JSON(http.StatusOK, "Healthy")
			} else {
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		})
	}
}
