package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/marisasha/ttl-check-app/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/marisasha/ttl-check-app/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		certificates := api.Group("/certificates")
		{
			certificates.GET("/", h.getAllCertificates)
			certificates.GET("/:id", h.getCertificate)
			certificates.POST("/add", h.addCertificate)
			certificates.POST("/check", h.checkCertificate)
			certificates.DELETE("/:id", h.deleteCertificate)
		}
	}
	return router
}
