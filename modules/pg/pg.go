package pg

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"saarm/pkg/models"
)


var DB *gorm.DB

// const (
//   DBHost = "postgres"
//   DBPort = "5400"
//   DBUsername = "tientran"
//   DBPassword = "tien123@"
//   DBName = "saarm_db"
//   DBTz = "Asia/Ho_Chi_Minh"
// )

func GetPgConnection() string {
  database := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
  os.Getenv("APP_ENV_POSTGRESQL_HOST"),
  os.Getenv("APP_ENV_POSTGRESQL_PORT"),
  os.Getenv("APP_ENV_POSTGRESQL_USERNAME"),
  os.Getenv("APP_ENV_POSTGRESQL_PASSWORD"),
  os.Getenv("APP_ENV_POSTGRESQL_DATABASE"),
)


  return database
}


func InitPg() *gorm.DB {
  var err error
  pgConn := GetPgConnection()
  DB, err = gorm.Open(postgres.Open(pgConn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
    QueryFields: true,
  })

  if err != nil {
    log.Panic(err)
  }

  log.Println("[InitPg] Connected!")
  log.Println("[InitPg] RUNNING MIGRATIONS....")

  DB.AutoMigrate(&models.User{}, &models.Room{})

  return DB

}
