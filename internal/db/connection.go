package db

import (
	"account-transaction-api/internal/config"
	"account-transaction-api/internal/db/seeders"
	"account-transaction-api/internal/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)

func Init(cfg *config.Config) *gorm.DB {

	var err error
	var db *gorm.DB

	log.Printf("initializing db connection for " + cfg.DB.Driver + "!")

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

	migration(db)

	// seeder
	userSeeder := seeders.NewOperationTypeSeeder(db)
	userSeeder.SetDefaultOperationTypes()

	return db
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&models.OperationType{})
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Transaction{})

	db.Model(&models.Transaction{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Transaction{}).AddForeignKey("operation_type_id", "operation_types(id)", "RESTRICT", "RESTRICT")

}
