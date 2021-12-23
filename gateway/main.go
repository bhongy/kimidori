package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/bhongy/kimidori/gateway/internal/api/v1"
)

func main() {
	router := gin.Default()
	api.Register(router)
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalln("Fail to start the service, err:", err)
	}
}
