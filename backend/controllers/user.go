package controllers

import (
	"net/http"

	"github.com/cerc-undip/tandera/types"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := new(types.User)
	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	users := []types.User{
		{
			ID:       1,
			Username: "user1",
			Password: "password",
		},
		{
			ID:       2,
			Username: "user2",
			Password: "password",
		},
	}

	return c.JSON(http.StatusOK, users)
}
