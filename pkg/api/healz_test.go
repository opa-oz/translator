package api_test

import (
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"translator/pkg/api"

	"github.com/gin-gonic/gin"
)

func TestHealzRoute(t *testing.T) {
	r := gin.Default()

	r.GET("/healz", api.Healz)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/healz", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	snaps.MatchSnapshot(t, w.Body.String())
}
