package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vosgaust/voicemod-challenge.git/internal/application/session"
	"github.com/vosgaust/voicemod-challenge.git/internal/application/user"
	"github.com/vosgaust/voicemod-challenge.git/internal/platform/auth"
	"github.com/vosgaust/voicemod-challenge.git/internal/platform/server"
	"github.com/vosgaust/voicemod-challenge.git/internal/platform/storage/mysql"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		log.Fatal(err)
	}

	jwtAuth := auth.NewJWTAuthentication(cfg.Auth.TimeToExpireDays, cfg.Auth.SignKey)

	userRepository := mysql.NewUserRepository(db)

	userService := user.NewUserService(userRepository)
	sessionService := session.NewSessionService(userService, jwtAuth)

	srv := server.New(cfg.Host, cfg.Port, userService, sessionService)

	if err := srv.Run(); err != nil {
		log.Fatal(srv.Run())
	}
}
