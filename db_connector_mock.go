package main

import "database/sql"

type DbConnectorMock struct {
	DbConnector *DbConnector
	Tx *sql.Tx
}

func NewDbConnectorMock(host, port, dbName, user, password string) *DbConnectorMock {
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
	mock := &DbConnectorMock{
		DbConnector: dbConnector,
	}
	return mock
}

func (c *DbConnectorMock) Init() {
	c.DbConnector.Init()
	c.Tx, _ = c.DbConnector.DB.Begin()
}

func (c *DbConnectorMock) Close() {
	c.Tx.Rollback()
	c.DbConnector.DB.Close()
}

func (c *DbConnectorMock) GetDb() IQuerier {
	return c.Tx
}
