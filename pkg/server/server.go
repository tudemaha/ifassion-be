package server

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tudemaha/ifassion-be/routes"
)

func StartServer() {
	log.Println("INFO StartServer: server is starting")

	router := gin.Default()
	routes.SetupRouterGroup(router)

	config := cors.Config{
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	router.Static("/pdf", "./files/pdf")

	port := ":" + os.Getenv("PORT")
	err := router.Run(port)
	if err != nil {
		log.Fatalf("ERROR StartServer fatal error: %v", err)
	}

	log.Println("INFO StartServer: server started successfully")
}
