package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/orvice/monitor-client/internal/config"
	"github.com/orvice/monitor-client/internal/hub"
	"net/http"
)

func Init() {
	r := gin.Default()
	r.GET("/", index)
	r.GET("/status", status)

	err := r.Run(config.WebAddr)
	if err != nil {
		fmt.Println("init web error", err)
	}
}

func index(c *gin.Context) {
	ni, err := hub.Monitor.GetInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, ni)
}

func status(c *gin.Context) {
	ni, err := hub.Monitor.GetInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
	}
	c.JSON(http.StatusOK, ni)
}
