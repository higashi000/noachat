package router

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	router := NewRouter()

	msgData := `{"text":"Test"}`

	req := httptest.NewRequest("POST", "/send", strings.NewReader(msgData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	router.E.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, msgData, rec.Body.String())
}
