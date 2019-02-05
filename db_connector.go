package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const dbDriver = "postgres"
const dbSSLMode = "disable"

type IQuerier interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type IDbConnector interface {
	Init()
	Close()
	GetDb() IQuerier
}

type dbConnectionData struct {
	host     string
	port     string
	dbName   string
	user     string
	password string
}

type DbConnector struct {
	dbConnectionData
	DB *sql.DB
}

func NewDbConnector(host, port, dbName, user, password string) IDbConnector {
	connectionData := dbConnectionData{
		host:     host,
		port:     port,
		dbName:   dbName,
		user:     user,
		password: password,
	}
	dbConnector := &DbConnector{
		dbConnectionData: connectionData,
	}
	return dbConnector
}

func (c *DbConnector) Init() {
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		c.host, c.port, c.dbName, c.user, c.password, dbSSLMode)
	var err error
	c.DB, err = sql.Open(dbDriver, connStr)
	if err != nil {
		panic(err)
	}
}

func (c *DbConnector) Close() {
	err := c.DB.Close()
	if err != nil {
		panic(err)
	}
}

func (c *DbConnector) GetDb() IQuerier {
	return c.DB
}
