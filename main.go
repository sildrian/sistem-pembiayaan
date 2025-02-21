package main

import (
	"sistem-pembiayaan/config"
	"sistem-pembiayaan/routes"
)

func main() {
	// Initialize database connection
	config.InitDB()
	defer config.DB.Close()

	// Initialize routes
	routes.Router()
}
