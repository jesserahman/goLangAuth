package domain

import (
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

	var result string
	err := a.dbClient.Select(&result, authQuery)
	if err != nil {
		logger.Error("Error querying Accounts table " + err.Error())
		return errs.NewUnexpectedError("unexpected db error")
	}

	fmt.Println("result is ", result)

	return nil
}

func NewAuthRepositoryDbConnection(dbClient *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{dbClient}
}
