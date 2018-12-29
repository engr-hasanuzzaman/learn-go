package main
import (
	"net/http"
	"github.com/labstack/echo"
	"handler"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "wellcome to echo world")
}

func showUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "User id" + id)
}

func allUser(c echo.Context) error {
	return e.Js
}

func main()  {
	e := echo.New()

	e.GET("/", index)
	// user routes
	e.POST("/users", createUser)
	e.GET("/users", handler.Index)
	e.GET("/users/:id", showUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1234"))
}

type User struct{
	Name string `json:"name"`
	Title int `json:"title"`
	Rattings []int `json:"rattings"`
}