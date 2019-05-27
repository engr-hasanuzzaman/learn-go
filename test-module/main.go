package main

import (
	_ "net/http"
	
	"github.com/labstack/echo"
	"test/handler"
)

func main() {
	e := echo.New()
	e.GET("/", handler.Index)
	e.GET("/hello", handler.Hello)
	e.Logger.Fatal(e.Start(":1323"))
}