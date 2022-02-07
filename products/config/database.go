package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBConnection interface {
	GetDB() (*sql.DB, error)
}

type DBConnectionImpl struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewDBConnection(
	host string,
	port string,
	username string,
	password string,
	database string,
) DBConnection {
	return &DBConnectionImpl{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
	}
}

func (c *DBConnectionImpl) GetDB() (*sql.DB, error) {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database)
	//dsnTwo := fmt.Sprintf("%s:%s@%s:5432/%s?sslmode=disable", c.Username, c.Password, c.Host, c.Database)
	//dsn := fmt.Sprintf("user=%s "+
	//	"password=%s "+
	//	"host=%s "+
	//	"port=%v "+
	//	"dbname=%s "+
	//	"sslmode=disable", c.Username, c.Password, c.Host, c.Port, c.Database)
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatalf("cannot open database: %v", err)
	}

	return db, err
}
