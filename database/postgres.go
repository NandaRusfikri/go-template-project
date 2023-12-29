package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-template-project/dto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDBPostgres(config dto.ConfigEnvironment) (*gorm.DB, error) {
	log.Debug("🔌 Connecting into Database")
	dbHost := config.DbHost
	dbUsername := config.DbUser
	dbPassword := config.DbPass
	dbName := config.DbName
	dbPort := config.DbPort
	dbSSLMode := config.DbSslmode
	timezone := config.Timezone

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
