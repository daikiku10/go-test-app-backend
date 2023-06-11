package handler

import (
	"net/http"

	"github.com/daikiku10/go-test-app-backend/context"
)

type postResponse struct {
	TestInt    int    `json:"testInt"`
	TestString string `json:"testString"`
}

func PostHandler(c *context.APIContext) error {
	response := postResponse{
		TestInt:    21,
		TestString: "postTestです。",
	}
	return c.JSON(http.StatusCreated, response)
}
