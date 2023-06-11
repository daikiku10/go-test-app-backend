package context

import "github.com/labstack/echo/v4"

type APIContext struct {
	echo.Context
}

func NewContext(c echo.Context) *APIContext {
	return &APIContext{
		Context: c,
	}
}

func Convert(h func(c *APIContext) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*APIContext))
	}
}
