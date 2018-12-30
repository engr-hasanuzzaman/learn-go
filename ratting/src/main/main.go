package main
import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "wellcome to echo world")
}

func main()  {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())

	e.GET("/", index)
	e.Logger.Fatal(e.Start(":5000"))
}
