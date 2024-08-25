package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
    "drugs/user"
)


func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

func getHome(c echo.Context) error {
	type User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	userD := User{
        ID:    1,
        Name:  "John Doe",
        Email: "john.doe@example.com",
    }
    return c.JSON(http.StatusOK, userD)
}

func main() {
	e := echo.New()
	e.GET("/show", show)
	e.GET("/", getHome)
	e.GET("/users/:id", user.GetUser)
	e.Logger.Fatal(e.Start(":8080"))
}
