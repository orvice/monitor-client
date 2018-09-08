package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func web() {
	r := gin.Default()
	r.GET("/", index)
	r.GET("/ws", func(c *gin.Context) {
		serveWs(h, c.Writer, c.Request)
	})
	r.Run(webAddr)
}

func index(c *gin.Context) {
	ni, err := mtr.GetInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, ni)
}
