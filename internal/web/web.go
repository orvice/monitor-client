package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/orvice/monitor-client/internal/config"
	"github.com/orvice/monitor-client/internal/hub"
	"github.com/orvice/monitor-client/pkg/vnstat"
	"net/http"
	"time"
)

func Init() {
	r := gin.Default()
	r.GET("/", index)
	r.GET("/status", status)
	r.GET("/vnstat/:interface", vnstatInterface)

	err := r.Run(config.WebAddr)
	if err != nil {
		fmt.Println("init web error", err)
	}
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"time":   time.Now().Unix(),
		"header": c.Request.Header,
	})
}

func status(c *gin.Context) {
	ni, err := hub.Monitor.GetInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
	}
	c.JSON(http.StatusOK, ni)
}

func vnstatInterface(c *gin.Context) {
	c.JSON(200, vnstat.VN(c.Param("interface")))
}
