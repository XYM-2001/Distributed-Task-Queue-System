package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load templates
	r.LoadHTMLGlob("internal/dashboard/templates/*")

	// Setup routes
	r.GET("/", handleDashboard)
	r.GET("/tasks", handleTasks)
	r.GET("/workers", handleWorkers)

	r.Run(":3000")
}

func handleDashboard(c *gin.Context) {
	// Get metrics and stats
	stats := getDashboardStats()
	c.HTML(http.StatusOK, "dashboard.html", stats)
}
