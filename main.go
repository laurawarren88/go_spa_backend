package main

import (
	"fmt"
	"log"

	"github.com/laurawarren88/go_spa_backend.git/config"
	"github.com/laurawarren88/go_spa_backend.git/database"
)

func init() {
	config.LoadEnv()
	config.SetGinMode()
}

func main() {
	database.ConnectToDB()
	db := database.GetDB()

	router := config.SetupServer()

	config.SetupHandlers(router, db)

	fmt.Printf("Starting the server on port %s\n", config.GetEnv("PORT", "8000"))
	if err := router.Run(":" + config.GetEnv("PORT", "8000")); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
