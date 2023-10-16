package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tudemaha/ifassion-be/routes"
)

func StartServer() {
	log.Println("INFO StartServer: server is starting")

	router := gin.Default()
	routes.SetupRouterGroup(router)

	err := router.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("ERROR StartServer fatal error: %v", err)
	}

	log.Println("INFO StartServer: server started successfully")
}
