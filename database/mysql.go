package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-template-project/dto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDBMysql(config dto.ConfigEnvironment) (*gorm.DB, error) {
	logrus.Debug("🔌 Connecting to Database")
	dbUsername := config.DB_USER
	dbPassword := config.DB_PASS
	dbName := config.DB_NAME
	dbHost := config.DB_HOST
	dbPort := config.DB_PORT
	//timezone := config.TIMEZONE

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		defer logrus.Errorln("❌ Error Connecting to Database: ", err.Error())
		return nil, err
	}

	err = MigrateDBSQL(db)
	if err != nil {
		logrus.Errorln("❌ Error Migrating Database: ", err.Error())
		return nil, err
	}

	logrus.Debug("✅ Connected to Database successfully")

	return db, nil
}
