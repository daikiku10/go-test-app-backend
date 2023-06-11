package handler

import (
	"net/http"

	"github.com/daikiku10/go-test-app-backend/context"
)

type APIGetResponse struct {
	TestInt    int    `json:"testInt"`
	TestString string `json:"testString"`
}

func APIGetHandler(c *context.APIContext) error {

	name := c.QueryParam("name")
	println(name)

	response := APIGetResponse{
		TestInt:    21,
		TestString: "apiGetTestです。",
	}
	return c.JSON(http.StatusOK, response)
}
