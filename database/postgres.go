package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-template-project/dto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBPostgres(config dto.ConfigDatabase, timezone string) (*gorm.DB, error) {
	log.Debug("🔌 Connecting into Database")
	dbHost := config.Host
	dbUsername := config.User
	dbPassword := config.Pass
	dbName := config.Name
	dbPort := config.Port
	dbSSLMode := config.SSLmode

	path := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		dbHost, dbUsername, dbPassword, dbName, dbPort, dbSSLMode, timezone)

	db, err := gorm.Open(postgres.Open(path), &gorm.Config{})

	if err != nil {
		log.Errorln("❌ Error Connect into Database ", err.Error())
		return nil, err
	}

	if err := migrateDB(config); err != nil {
		return nil, err
	}

	if err = createSeeds(db); err != nil {
		return nil, err
	}

	log.Debug("🔌 Connect into Database Success")

	return db, nil
}
