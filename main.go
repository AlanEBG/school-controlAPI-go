package main

import (
	"fmt"
	"log"

	"github.com/AlanEBG/school-controlAPI-go/database"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontro archivo .env, usando valores por defecto")
	}

	// Conectar a la base de datos
	database.Connect()

	fmt.Println("Sistema de Control Escolar iniciado correctamente")
}
