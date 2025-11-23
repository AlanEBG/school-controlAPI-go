package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AlanEBG/school-controlAPI-go/database"
	"github.com/AlanEBG/school-controlAPI-go/router"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontro archivo .env, usando valores por defecto")
	}

	// Conectar a la base de datos
	database.Connect()

	// Configurar router
	r := router.SetupRouter()

	// Obtener puerto desde variable de entorno o usar default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	fmt.Printf("Servidor corriendo en http://localhost:%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
