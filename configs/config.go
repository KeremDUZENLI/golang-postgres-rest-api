package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func DatabaseDB() {
	dsn := "host=localhost user=test password=test dbname=postgres_test port=5000 sslmode=disable"

	database_postgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("NOT CONNECTED")
	}

	Database = database_postgres
}
