package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"

	"github.com/bnock/nockchat-api-go/models"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Env struct {
	users    models.UserModel
	channels models.ChannelModel
}

func main() {
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

	env := &Env{
		channels: models.ChannelModel{DB: db},
		users:    models.UserModel{DB: db},
	}

	e := echo.New()

	// Public routes
	public := e.Group("/")
	public.GET("/", env.index).Name = "index"

	// Protected routes
	protected := e.Group("/p")
	protected.Use(echojwt.JWT([]byte(os.Getenv("APP_SECRET"))))

	protected.GET("/channels/:channel", env.channel).Name = "channels"

	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}

func (env *Env) index(c echo.Context) error {
	users, err := env.users.All()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to fetch users")
	}

	return c.JSON(http.StatusOK, users)
}

func (env *Env) channel(c echo.Context) error {
	userId, err := UserIDFromCtx(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to fetch user id")
	}

	user, err := env.users.UserById(userId)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to fetch user")
	}

	channel, err := env.channels.ChannelById(c.Param("channel"))
	if err != nil {
		return c.String(http.StatusNotFound, "Unable to fetch channel")
	}

	if user.ID != channel.OwnerID && !slices.Contains(channel.MemberIDs, user.ID) {
		return c.String(http.StatusForbidden, "You do not have access to this channel")
	}

	return c.JSON(http.StatusOK, channel)
}
