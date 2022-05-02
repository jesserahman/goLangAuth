package app

import (
	"fmt"
	"github.com/jesserahman/goLangAuth/domain"
	"github.com/jesserahman/goLangAuth/logger"
	"github.com/jesserahman/goLangAuth/service"
	"log"
	"net/http"
	"os"
	"time"

	mux2 "github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func sanityCheck() {
	if os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Env variables not defined...")
	}
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sanityCheck()
	dbClient := getDbClient()

	authHandler := AuthHandler{service.NewAuthService(domain.NewAuthRepositoryDbConnection(dbClient), domain.GetRolePermissions())}
	registerHandler := RegisterHandler{service.NewRegisterService(domain.NewRegisterRepositoryDbConnection(dbClient))}

	router := mux2.NewRouter()
	router.HandleFunc("/auth/login", authHandler.handleLogin).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", authHandler.handleVerify).Methods(http.MethodGet)
	router.HandleFunc("/register", registerHandler.handleRegister).Methods(http.MethodPost)

	port := os.Getenv("SERVER_PORT")

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
	dbClient, err := sqlx.Open("mysql", datasource)
	if err != nil {
		logger.Error("Error connecting to the DB " + err.Error())
		panic(err)
	}
	// See "Important settings" section.
	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)
	return dbClient
}
