package user

import (
	"net/http"
    "github.com/labstack/echo/v4"
	"fmt"
    "strconv"
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