package test

import (
	entity "countries-api/entity"
	"countries-api/http/gin/handler"
	"countries-api/http/gin/server"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetCountriesByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDbInterface(ctrl)
	handler := handler.Handler{Db: m}
	s := server.Server{Handler: &handler}
	router := s.SetupRouter()

	data := &entity.Country{
		ID:            "1",
		Name:          "nigeria",
		ShortName:     "NGN",
		Continent:     "Africa",
		IsOperational: true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	m.EXPECT().Find(data.ID).Return(data, nil)
	jsonFile, err := json.Marshal(data)
	if err != nil {
		t.Error("Failed to marshal file")
	}

	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/countries/1", nil)

	router.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)
	assert.Contains(t, rw.Body.String(), string(jsonFile))
}
