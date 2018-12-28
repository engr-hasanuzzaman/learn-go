package main
import (
	"net/http"
	"github.com/labstack/echo"
	"handler"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "wellcome to echo world")
}

func main()  {
	e := echo.New()

	e.GET("/", index)
	e.GET("/users", handler.Index)
	e.Logger.Fatal(e.Start(":1234"))
}