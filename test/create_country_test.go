package test

import (
	"countries-api/http/gin/handler"
	"countries-api/http/gin/server"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDbInterface(ctrl)
	handler := handler.Handler{Db: m}
	s := server.Server{Handler: &handler}
	s.SetupRouter()

	// Does not make any assertions. Executes the anonymous functions and returns
	// its result when Bar is invoked with 99.
	m.EXPECT().Create(gomock.Any()).Return(gomock.Any())

}