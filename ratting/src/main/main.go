package main
import (
	"net/http"
	"github.com/labstack/echo"
	_ "handler"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "wellcome to echo world")
}

func main()  {
	e := echo.New()

	e.GET("/", index)
	// user routes
	e.POST("/users", createUser)
	e.GET("/users", allUser)
	e.GET("/users/:id", showUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.GET("/users/search", searchUser)

	// 
	e.Logger.Fatal(e.Start(":5000"))
}

type User struct{
	Name string `json:"name"`
	Title int `json:"title"`
	Rattings []int `json:"rattings"`
}

// user route handler
func showUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "User id " + id + "\n")
}

func allUser(c echo.Context) error {
	return c.String(http.StatusOK, "show all user list\n")
}

func updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Update user " + c.Param("id") + "\n")
}

func searchUser(c echo.Context) error {
	// title := c.QueryParam("title")
	return c.String(http.StatusOK, "passing search tearm are " + c.QueryParam("title") + "\n")
}

func createUser(c echo.Context) error {
	return c.String(http.StatusOK, "New user created\n")
}

func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "user updated " + c.Param("id") + "\n")
}