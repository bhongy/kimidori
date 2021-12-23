// Package service ...
package service

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/bhongy/kimidori/gateway/internal/service/repository"
	"github.com/gin-gonic/gin"
)

type getServiceResponseBody struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
}

func GetService(c *gin.Context) {
	repo := repository.NewFileSystem("config/services")
	be, err := repo.ByServiceName(c.Param("svc"))
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.AbortWithError(http.StatusBadRequest, err)
		} else {
			log.Fatalln(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	u, err := url.Parse(be.Origin)
	if err != nil {
		log.Fatalln(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	// proxy := httputil.NewSingleHostReverseProxy(target)
	// proxy.ServeHTTP(c.Writer, c.Request)

	body := getServiceResponseBody{
		Protocol: u.Scheme,
		Host:     u.Host,
	}
	c.JSON(http.StatusOK, body)
}
