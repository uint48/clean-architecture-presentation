package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type ConnInfo struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DBName string
}

func NewMySQLConnection(c *ConnInfo) (*sqlx.DB, error) {
	dbHost := c.Host
	dbPort := c.Port
	dbUser := c.User
	dbPass := c.Pass
	dbName := c.DBName

	// example of connection string: "test:test@(localhost:3306)/test"

	db, err := sqlx.Connect("mysql", dbUser+":"+dbPass+"@("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseMySQLConnection(client *sqlx.DB) {
	if client == nil {
		return
	}

	err := client.Close()
	if err != nil {
		return
	}

	log.Println("Connection to MySQL closed.")
}
