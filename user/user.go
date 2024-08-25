package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
	}
	user := User{
		ID:    i,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}
	return c.JSON(http.StatusOK, user)
}
