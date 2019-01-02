package main
import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"html/template"
	"io"

	"github.com/engr-hasanuzzaman/test/handler"
)

func index(c echo.Context) error {
	// return c.String(http.StatusOK, "wellcome to echo world")
	return c.Render(http.StatusOK, "hello", map[string]interface{}{
		"title": "hello",
		"msg": "Hello from index",
	})
}

func main()  {
	e := echo.New()

	t := &Template{
    templates: template.Must(template.ParseGlob("public/views/*.html")),
  }
	e.Renderer = t
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.Index)
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

// renderer
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// user model
type User struct{
	Name string `json:"name" xml:"name" query:"name" form:"name"`
	Title string `json:"title" xml:"title" query:"title" form:"title"`
	Rattings []int `json:"rattings" xml:"rattings" query:"rattings" form:"rattings"`
}

// user route handler
func showUser(c echo.Context) error {
	log.Info("this is info log")
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
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "user updated " + c.Param("id") + "\n")
}