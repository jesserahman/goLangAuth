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

func (a AuthRepositoryDb) CheckCredentials(request dto.NewAuthRequest) *errs.AppError {
	var authQuery = fmt.Sprintf("select * from users u where username = '%s' AND password = '%s'", request.Username, request.Password)

	var resultUser User
	err := a.dbClient.Get(&resultUser, authQuery)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error querying Accounts table " + err.Error())
		return errs.NewUnexpectedError("unexpected db error")
	}

	if err == sql.ErrNoRows {
		logger.Error("Error querying Accounts table " + err.Error())
		return errs.NewNotFoundError("username and password not found")
	}

	fmt.Println("result is ", resultUser)

	return nil
}

func NewAuthRepositoryDbConnection(dbClient *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{dbClient}
}
