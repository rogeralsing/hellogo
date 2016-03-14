package person

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestPersonRepository struct {
	people map[string]Person
}

func (this TestPersonRepository) GetPerson(id string) (Person, error) {
	val, ok := this.people[id]
	if ok {
		return val, nil
	}
	return val, fmt.Errorf("Person not found")
}

func (this TestPersonRepository) SavePerson(id string, person Person) error {
	return nil
}

func TestGetPerson(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	testPersonDB := TestPersonRepository{people: map[string]Person{"roger": Person{Name: "Roger", Age: 40}}}
	CreatePersonService(router, testPersonDB)
	req, _ := http.NewRequest("GET", "/api/v1/person/roger", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"name\":\"Roger\",\"age\":40,\"children\":null}\n", resp.Body.String())
}
