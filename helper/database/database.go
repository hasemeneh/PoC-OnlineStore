package database

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"log"
)

type logger func(v ...interface{})

type DatabaseConnector struct {
	fatalln logger
}

func New() *DatabaseConnector {
	return &DatabaseConnector{
		fatalln: log.Fatalln,
	}
}

type ExecFunc func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

const MySQLDatabaseDriver = "mysql"

func (d *DatabaseConnector) Connect(connectionString string) *sqlx.DB {
	db, err := sqlx.Open(MySQLDatabaseDriver, connectionString)

	if err != nil {
		d.fatalln("error Opening Database", err)
	}
	return db
}
