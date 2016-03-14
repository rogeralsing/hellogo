package person

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPerson(t *testing.T) {
	router := gin.Default()
	CreatePersonService(router, nil)
	req, _ := http.NewRequest("GET", "/api/v1/person/roger", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "bar")
}
