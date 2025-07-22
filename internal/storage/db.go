package storage

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dialector := sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        "database/test.db",
	}

	DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	err = DB.AutoMigrate(&Request{})
	if err != nil {
		log.Fatalf("Ошибка миграции таблиц: %v", err)
	}
}
