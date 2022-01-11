package domain

import "github.com/jmoiron/sqlx"

type AuthRepositoryDb struct {
	dbClient *sqlx.DB
}

func NewAuthRepositoryDbConnection(dbClient *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{dbClient}
}
