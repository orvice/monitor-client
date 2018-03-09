package main

import (
	"github.com/gin-gonic/gin"
)

func web() {
	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		serveWs(h, c.Writer, c.Request)
	})
	r.Run(webAddr)
}
