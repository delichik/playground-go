package main

import (
	"github.com/labstack/echo/v4"

	"router/route"
)

func main() {
	e := echo.New()
	route.InitRouter(e.Group("/"))
}
