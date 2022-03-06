package domain

import (
	"database/sql"
	"fmt"
	"goLangAuth/dto"
	"goLangAuth/errs"
	"goLangAuth/logger"

	"github.com/jmoiron/sqlx"
)

type AuthRepositoryDb struct {
	dbClient *sqlx.DB
}

func (a AuthRepositoryDb) CheckCredentials(request dto.NewAuthRequest) (*User, *errs.AppError) {
	var authQuery = fmt.Sprintf("select username, u.customer_id, role,  GROUP_CONCAT(a.account_id) as account_ids from users u LEFT JOIN accounts a ON a.customer_id = u.customer_id where username = '%s' and password = '%s'", request.Username, request.Password)

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

	if resultUser.Role != "admin" && resultUser.Role != "user" {
		return nil, errs.NewUnexpectedError("user has invalid role")
	}
	return &resultUser, nil
}

func NewAuthRepositoryDbConnection(dbClient *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{dbClient}
}
