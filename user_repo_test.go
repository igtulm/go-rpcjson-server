package main

import (
	"database/sql"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}

type fixture struct {
	t *testing.T
	tx  *sql.Tx
	db IDbConnector
	repo IUserRepository
}

func tearUp(t *testing.T) *fixture {
	cfg := NewConfig()
	cfg.Init()

	db := NewDbConnector(cfg.DB.Host, cfg.DB.Port, cfg.DB.Database, cfg.DB.User, cfg.DB.Password)
	db.Init()

	tx, err := db.GetDb().Begin()
	if err != nil {
		require.NoError(t, err)
	}

	repo := NewUserRepository(db)

	return &fixture{
		t: t,
		db:  db,
		tx: tx,
		repo: repo,
	}
}

func (fx *fixture) tearDown() {
	defer fx.db.Close()
	fx.tx.Rollback()
}

func (fx *fixture) createTestUser(login string) *User {
	now := time.Now()
	var ID string
	row := fx.tx.QueryRow(`INSERT INTO users(login, created_at) VALUES($1, $2) RETURNING id`,
		login,
		now)
	err := row.Scan(&ID)
	if err != nil {
		require.NoError(fx.t, err)
	}
	user := &User{
		ID: ID,
		Login: login,
		CreatedAt: now.Unix(),
	}
	return user
}

func TestUserRepository_GetByLogin(t *testing.T) {
	t.Run("Should be albe to get user by his login", func(t *testing.T) {
		fx := tearUp(t)
		defer fx.tearDown()

		login := RandString(20)
		expUser := fx.createTestUser(login)

		actUser, err := fx.repo.GetByLogin(login)
		require.NoError(fx.t, err)
		require.Equal(fx.t, expUser, actUser)
	})
}
