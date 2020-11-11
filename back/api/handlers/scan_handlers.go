package handlers

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/scan"
	"github.com/gin-gonic/gin"
)

func FullScanHandler(c *gin.Context) {
	go scan.StartFullScan()
	c.Done()
}
