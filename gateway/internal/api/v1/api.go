package api

import (
	"github.com/gin-gonic/gin"

	"github.com/bhongy/kimidori/gateway/internal/service"
)

func Register(router *gin.Engine) {
	router.GET("/services/:svc", service.GetService)
}
