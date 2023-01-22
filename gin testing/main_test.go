package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestApiMain(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/1/", nil)

	r := gin.Default()
	registerRoutes(r)
	r.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Fail()
	}
	t.Log(res.Body)

}
