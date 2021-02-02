package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	userJSON := `
		{
			"username": "testusername",
			"email": "test@example.com"
		}
	`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	assert.NoError(t, CreateUser(c))
	assert.Equal(t, http.StatusCreated, rec.Code)
	t.Log(rec.Body.String())
}

func TestGetUsers(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	assert.NoError(t, GetUsers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}
