package test

import (
	"countries-api/entity"
	"countries-api/http/gin/handler"
	"countries-api/http/gin/server"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCountries(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDbInterface(ctrl)
	handler := handler.Handler{Db: m}
	s := server.Server{Handler: &handler}
	router := s.SetupRouter()

	payload := &struct {
		Name          string `json:"name" bson:"name" binding:"required"`
		ShortName     string `json:"short_name" bson:"short_name" binding:"required"`
		Continent     string `json:"continent" bson:"continent" binding:"required"`
		IsOperational bool   `json:"is_operational" bson:"is_operational" binding:"required"`
	}{
		Name:          "nigeria",
		ShortName:     "NGN",
		Continent:     "Africa",
		IsOperational: true,
	}

	data := &entity.Country{
		Name:          "nigeria",
		ShortName:     "NGN",
		Continent:     "Africa",
		IsOperational: true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	m.EXPECT().Find(gomock.Any()).Return(data, nil)
	m.EXPECT().Update(*data, gomock.Any()).Return(nil, nil)

	jsonFile, err := json.Marshal(payload)
	if err != nil {
		t.Error("Failed to marshal file")
	}
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/countries/1", strings.NewReader(string(jsonFile)))

	router.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)
	assert.Contains(t, rw.Body.String(), "nigeria")

	// Does not make any assertions. Returns 103 when Bar is invoked with 101.

}
