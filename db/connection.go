package db

import (
	"account-transaction-api/config"
	"account-transaction-api/db/seeders"
	"account-transaction-api/models"
	"fmt"
	"go.uber.org/zap"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB  *gorm.DB
	err error
)

func Init(cfg *config.Config) *gorm.DB {

	var db = DB

	zap.L().Info("initializing db connection for " + cfg.DB.Driver + "!")

	if cfg.DB.Driver == "postgres" { // POSTGRES
		db, err = gorm.Open("postgres",
			fmt.Sprintf("host=%s port=%s user=%s dbname=%s  sslmode=disable password=%s",
				cfg.DB.Host,
				cfg.DB.Port,
				cfg.DB.User,
				cfg.DB.Name,
				cfg.DB.Password))
		if err != nil {
			panic(err)
		}
	} else {
		panic(fmt.Sprintf("db driver %s not supported", cfg.DB.Driver))
	}

	db.LogMode(cfg.DB.LogMode)
	db.DB().SetMaxIdleConns(cfg.DB.MaxIdleConnections)
	db.DB().SetMaxOpenConns(cfg.DB.MaxConnections)
	db.DB().SetConnMaxLifetime(time.Duration(cfg.DB.MaxLifetime) * time.Second)

	// seeder
	userSeeder := seeders.NewOperationTypeSeeder(db)
	userSeeder.SetDefaultOperationTypes()

	migration()

	return db
}

func migration() {
	DB.AutoMigrate(&models.OperationType{})
	DB.AutoMigrate(&models.Account{})
	DB.AutoMigrate(&models.Transaction{})
}
