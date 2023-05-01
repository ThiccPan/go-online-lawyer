package config

import (
	"fmt"
	"log"
	"os"

	"github.com/thiccpan/go-online-lawyer/entities"
	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DB_Host, config.DB_Username, config.DB_Password, config.DB_Name, config.DB_Port)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db.AutoMigrate(
		&entities.Pengacara{},
	)
	
	return db
}
