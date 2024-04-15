package pg

import (
  "fmt"
  "log"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"

  "saarm/pkg/models"
)


var DB *gorm.DB

var (
  DBHost = "127.0.0.1"
  DBPort = "5401"
  DBUsername = "tientran"
  DBPassword = "tien123@"
  DBName = "saarm_db"
  DBTz = "Asia/Ho_Chi_Minh"
)

func GetPgConnection() string {
  database := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s sslmode=disable",
  DBHost,
  DBPort,
  DBUsername,
  DBPassword,
  DBName,
  DBTz,
)


  return database
}


func InitPg() *gorm.DB {
  var err error
  connection := GetPgConnection()
  DB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
  if err != nil {
    log.Panic(err)
  }

  log.Println("[InitPg] Connected!")
  log.Println("[InitPg] RUNNING MIGRATIONS....")

  DB.AutoMigrate(&models.User{}, &models.Room{})

  return DB

}
