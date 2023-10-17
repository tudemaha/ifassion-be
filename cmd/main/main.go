package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/tudemaha/ifassion-be/pkg/server"
)

func main() {
	server.StartServer()
}
