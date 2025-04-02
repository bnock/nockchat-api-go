package cmd

import (
	"database/sql"
	"log"
	"os"

	"github.com/bnock/nockchat-api-go/internal/handlers"
	"github.com/bnock/nockchat-api-go/internal/repositories"
	"github.com/bnock/nockchat-api-go/internal/server"
	"github.com/bnock/nockchat-api-go/internal/services"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Serve() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbCfg := mysql.Config{
		Addr:      os.Getenv("MYSQL_HOST"),
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		DBName:    os.Getenv("MYSQL_DATABASE"),
		Net:       "tcp",
		ParseTime: true,
		Params:    map[string]string{"charset": "utf8mb4"},
	}

	db, err := sql.Open("mysql", dbCfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	r := repositories.NewRepositories(db)
	s := services.NewServices(r)
	h := handlers.NewHandlers(s)

	srv := server.NewServer(
		server.WithRoutes(h),
	)

	srv.Run()
}
