package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

type User struct {
	ID    int    `json:"id" validate:"gte=0"`
	Name  string `json:"name" validate:"required,min=3,max=32"`
	Email string `json:"email" validate:"required,email"`
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
	if err := validate.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make(map[string]string)
		for _, fieldErr := range validationErrors {
			errors[fieldErr.Field()] = fieldErr.Tag()
		}
		return c.JSON(http.StatusBadRequest, errors)
	}
	return c.JSON(http.StatusOK, user)
}
