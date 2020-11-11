package handlers

import (
	"github.com/gin-gonic/gin"
)

func StatusHandler(c *gin.Context) {
	c.Done()
}
