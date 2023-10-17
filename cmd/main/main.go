package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/tudemaha/ifassion-be/internal/export/dto"
	"github.com/tudemaha/ifassion-be/internal/export/services"
)

func main() {
	// server.StartServer()
	services.CreatePdf(dto.ResponseData{
		ID:         "652cd4a638b3c8928ca8b52a",
		Passion:    "Data Analyst",
		Time:       "2023-10-17 12:12:12",
		Indicators: []string{"suka belajar backend", "suka belajar frontend", "sdlhfsdoihfsi"},
	})
}
