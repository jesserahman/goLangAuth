package app

import (
	"fmt"
	mux2 "github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"goLangAuth/domain"
	"goLangAuth/logger"
	"goLangAuth/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
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

	authHandler := AuthHandler{service.NewAuthService(domain.NewAuthRepositoryDbConnection(dbClient))}
	registerHandler := RegisterHandler{service.NewRegisterService(domain.NewRegisterRepositoryDbConnection(dbClient))}

	router := mux2.NewRouter()
	router.HandleFunc("/login", authHandler.handleLogin).Methods(http.MethodPost)
	router.HandleFunc("/register", registerHandler.handleRegister).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)
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
