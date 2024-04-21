package api_test

import (
	"bytes"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"translator/pkg/api"
	"translator/pkg/config"
	"translator/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func TestGetRoute(t *testing.T) {
	r := gin.Default()
	cfg, _ := config.GetConfig()
	r.Use(middlewares.CfgMiddleware(cfg))

	r.POST("/translate/:fromLang/:toLang", api.Translate)

	t.Run("should translate text", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/translate/ru/en", bytes.NewBuffer([]byte("Привет мир!")))
		r.ServeHTTP(w, req)

		snaps.MatchSnapshot(t, w.Body.String())
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
