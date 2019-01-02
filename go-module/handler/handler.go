package handler
import (
	"net/http"
	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "test handler")
}

