package database

import (
	"fmt"
	"log"
	"os"

	"github.com/AlanEBG/school-controlAPI-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect inicializa la conexión a la base de datos
func Connect() {
	var err error

	// Obtener la ruta de la BD desde variable de entorno o usar default
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "school.db"
	}

	// Configurar GORM con logger
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	fmt.Println("Conexion a la base de datos exitosa")

	// Ejecutar migraciones automáticas
	err = DB.AutoMigrate(
		&models.Student{},
		&models.Subject{},
		&models.Grade{},
	)

	if err != nil {
		log.Fatal("Error al ejecutar migraciones:", err)
	}

	fmt.Println("Migraciones ejecutadas correctamente")
}

// GetDB retorna la instancia de la base de datos
func GetDB() *gorm.DB {
	return DB
}
