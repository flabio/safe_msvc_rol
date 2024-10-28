package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/msvc_rol/infrastructure/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB
var dbOnce sync.Once

func LoadEnv() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println("Error cargando el archivo .env", errEnv)
	}
}
func GetDatabaseInstance() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		dbInstance, err = DatabaseConnection()
		if err != nil {
			log.Fatalf("Error al inicializar la base de datos: %v", err)
		}
	})
	return dbInstance
}

func DatabaseConnection() (*gorm.DB, error) {
	LoadEnv()
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  sslmode=%s ",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
		DB_SSLMODE,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar a la base de datos: %w", err)
	}
	err = db.AutoMigrate(&entities.Rol{})
	if err != nil {
		return nil, fmt.Errorf("no se pudo migrar la base de datos: %w", err)
	}
	return db, nil
}

func CloseConnection() {
	db := GetDatabaseInstance()
	dbSQL, err := db.DB()
	if err != nil {
		log.Println(err.Error())
	}
	err = dbSQL.Close()
	if err != nil {
		log.Println("Error closing database connection", err)
	}
}
