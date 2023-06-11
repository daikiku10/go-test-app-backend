package handler

import (
	"net/http"

	"github.com/daikiku10/go-test-app-backend/context"
)

type getResponse struct {
	TestInt    int    `json:"testInt"`
	TestString string `json:"testString"`
}

func GetHandler(c *context.APIContext) error {

	name := c.QueryParam("name")
	println(name)

	response := getResponse{
		TestInt:    21,
		TestString: "getTestです。",
	}
	return c.JSON(http.StatusOK, response)
}
