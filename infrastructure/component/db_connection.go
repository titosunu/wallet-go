package component

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/titosunu/wallet-go/infrastructure/config"
)

func GetDatabaseConnection(config *config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name)

	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error when open connection %s", err.Error())
	}

	err = connection.Ping()
	if err != nil {
		log.Fatalf("error when open connection %s", err.Error())
	}

	return connection
}