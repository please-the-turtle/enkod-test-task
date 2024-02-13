package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/please-the-turtle/encod-test-task/http"
	"github.com/please-the-turtle/encod-test-task/logic"
	"github.com/please-the-turtle/encod-test-task/postgresql"
	log "github.com/sirupsen/logrus"
)

func main() {
	port := os.Getenv("APP_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	log.Info(dsn)

	conn, err := dbr.Open("postgres", dsn, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	session := conn.NewSession(nil)
	personRepo := postgresql.NewPostgresPersonRepository(session)

	timeout := time.Duration(time.Second * 60)
	personLogic := logic.NewPersonLogic(personRepo, timeout)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	http.NewPersonHandler(e, personLogic)

	e.Logger.Fatal(e.Start(":" + port))
}
