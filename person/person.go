package person

import (
	"github.com/gin-gonic/gin"
	"github.com/rhinoman/couchdb-go"
	"net/http"
)

type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Children []string `json:"children"`
}

func CreatePersonService(router *gin.Engine, db *couchdb.Database) {
	person := router.Group("api/v1/person")
	{
		person.GET(":id", func(c *gin.Context) {
			id := c.Param("id")
			var doc Person
			if _, err := db.Read(id, &doc, nil); err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			c.JSON(http.StatusOK, doc)
		})

		person.PUT(":id", func(c *gin.Context) {
			id := c.Param("id")
			var doc Person
			if err := c.BindJSON(&doc); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}
			rev, err := db.Read(id, &doc, nil)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			if _, err := db.Save(doc, id, rev); err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			c.JSON(http.StatusOK, doc)
		})
	}
}
