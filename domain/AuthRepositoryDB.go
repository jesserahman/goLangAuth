package domain

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"goLangAuth/dto"
	"goLangAuth/errs"
	"goLangAuth/logger"
)

type AuthRepositoryDb struct {
	dbClient *sqlx.DB
}

func (a AuthRepositoryDb) CheckCredentials(request dto.NewAuthRequest) (*User, *errs.AppError) {
	var authQuery = fmt.Sprintf("select * from users u where username = '%s' AND password = '%s'", request.Username, request.Password)

	var resultUser User
	err := a.dbClient.Get(&resultUser, authQuery)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error querying User table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected db error")
	}

	if err == sql.ErrNoRows {
		logger.Error("No rows found matching the username and password " + err.Error())
		return nil, errs.NewUnauthorizedError("username or password in incorrect")
	}
	return &resultUser, nil
}

func NewAuthRepositoryDbConnection(dbClient *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{dbClient}
}
