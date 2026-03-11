package app

import (
	"log"
	"os"
	"time"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(user, host, password, port, db string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	//  function auto migrate, create and generate schema table
	err = database.AutoMigrate(
		&domain.User{},
		&domain.Category{},
		&domain.Supplier{},
		&domain.Product{},
		&domain.StockIn{},
		&domain.StockOut{},
	)
	if err != nil {
		panic("failed to auto migrate schema")

	}

	return database
}
