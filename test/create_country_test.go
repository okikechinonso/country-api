package test

import (
	"countries-api/http/gin/handler"
	"countries-api/http/gin/server"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDbInterface(ctrl)
	handler := handler.Handler{Db: m}
	s := server.Server{Handler: &handler}
	router := s.SetupRouter()

	// Does not make any assertions. Executes the anonymous functions and returns
	// its result when Bar is invoked with 99.
	payload := &struct {
		Name          string `json:"name" bson:"name" binding:"required"`
		ShortName     string `json:"short_name" bson:"short_name" binding:"required"`
		Continent     string `json:"continent" bson:"continent" binding:"required"`
		IsOperational bool   `json:"is_operational" bson:"is_operational" binding:"required"`
	}{
		Name:          "Nigeria",
		ShortName:     "NGN",
		Continent:     "Africa",
		IsOperational: true,
	}

	m.EXPECT().Create(gomock.Any()).Return(nil,nil)

	jsonFile, err := json.Marshal(payload)
	if err != nil {
		t.Error("Failed to marshal file")
	}
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/countries", strings.NewReader(string(jsonFile)))

	router.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)

}
