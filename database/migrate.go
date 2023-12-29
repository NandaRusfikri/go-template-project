package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"go-template-project/dto"
)

func migrateDB(config dto.ConfigEnvironment) error {
	dbHost := config.DbHost
	dbUsername := config.DbUser
	dbPassword := config.DbPass
	dbName := config.DbName
	dbPort := config.DbPort
	dbSSLMode := config.DbSslmode

	m, err := migrate.New(
		"file://database/migrations",
		fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", dbUsername, dbPassword, dbHost, dbPort, dbName, dbSSLMode))
	if err != nil {
		log.Println("DbMigrate ", err)
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Println("DbMigrate ", err)
		return err
	}
	return nil

	// export POSTGRESQL_URL='postgres://nanda:nanda@localhost:5432/go-template-project?sslmode=disable'
	// migrate create -ext sql -dir database/migrations -seq create_table_user

	// migrate -database ${POSTGRESQL_URL} -path database/migrations up
	// migrate -database ${POSTGRESQL_URL} -path database/migrations down 2
	// migrate -database ${POSTGRESQL_URL} -path database/migrations version
	// migrate -database ${POSTGRESQL_URL} -path database/migrations force 2

}
