package handler

import (
	"github.com/alphablaster/mocker/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")

	{
		auth.POST("/sign_up", h.signUp)
		auth.POST("/sign_in", h.signIn)
	}

	mocks := router.Group("/mocks")

	{
		mocks.POST("/", h.createMock)
		mocks.GET("/", h.getAllMocks)
		mocks.GET("/:id", h.getMockById)
		mocks.PUT("/:id", h.updateMock)
		mocks.DELETE("/:id", h.deleteMock)
	}

	return router
}
