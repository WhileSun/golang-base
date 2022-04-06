package gdb

import (
	"fmt"

	"gorm.io/driver/postgres"
)

func postGresInit() {
	dbURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.User,
		databaseConfig.Name,
		databaseConfig.Password)
	dialector = postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	})
}
