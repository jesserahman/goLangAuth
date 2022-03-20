package domain

import (
	"fmt"
	"github.com/jesserahman/goLangAuth/dto"
	"github.com/jesserahman/goLangAuth/errs"

	"github.com/jmoiron/sqlx"
)

type RegisterRepositoryDb struct {
	dbClient *sqlx.DB
}

func (r RegisterRepositoryDb) RegisterNewUser(request dto.NewRegisterRequest) *errs.AppError {
	fmt.Println("inside register new user")
	return nil
}

func NewRegisterRepositoryDbConnection(dbClient *sqlx.DB) RegisterRepositoryDb {
	return RegisterRepositoryDb{dbClient}
}
