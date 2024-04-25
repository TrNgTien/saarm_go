package pg

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"saarm/pkg/models"
	"saarm/pkg/utilities"
)

var DB *gorm.DB

func GetPgConnection() string {
	database := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		utilities.GetValueOrDefault(os.Getenv("APP_ENV_POSTGRESQL_HOST"), "localhost"),
		utilities.GetValueOrDefault(os.Getenv("APP_ENV_POSTGRESQL_PORT"), "5400"),
		utilities.GetValueOrDefault(os.Getenv("APP_ENV_POSTGRESQL_USERNAME"), "tientran"),
		utilities.GetValueOrDefault(os.Getenv("APP_ENV_POSTGRESQL_PASSWORD"), "tien123@"),
		utilities.GetValueOrDefault(os.Getenv("APP_ENV_POSTGRESQL_DATABASE"), "saarm_db"),
	)

	return database
}

func InitPg() *gorm.DB {
	var err error

	pgConn := GetPgConnection()
	DB, err = gorm.Open(postgres.Open(pgConn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		QueryFields: true,
	})

	if err != nil {
		log.Panic(err)
	}

	log.Println("[InitPg] Connected!")
	log.Println("[InitPg] RUNNING MIGRATIONS....")

	DB.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.Account{},
		&models.Apartment{},
		&models.MetaLink{},
		&models.Role{},
		&models.UserRole{},
		&models.Permission{},
		&models.PermissionMapping{},
		&models.UserApartment{},
	)

	return DB

}
