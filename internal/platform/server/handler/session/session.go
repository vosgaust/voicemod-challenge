package session

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vosgaust/voicemod-challenge.git/internal/application/session"
)

type loginPayload struct {
	Email    string
	Password string
}

type response struct {
	Status string `json:"status"`
	Msg    string `json:"msg,omitempty"`
}

func LoginHandler(sessionService session.SessionService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		login := loginPayload{}

		if err := ctx.BindJSON(&login); err != nil {
			log.Infof("failed binding user: %v", err)
			ctx.JSON(http.StatusUnprocessableEntity, response{"error", "Login payload is incorrect"})
			return
		}

		token, err := sessionService.Authenticate(ctx, login.Email, login.Password)
		if err != nil {
			log.Infof("Failed authenticating user: %v", err)
			switch {
			case errors.Is(err, session.ErrIncorrectPassword), errors.Is(err, session.ErrUserNotFound):
				ctx.JSON(http.StatusBadRequest, response{"error", "invalid password"})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, response{"error", ""})
				return
			}
		}

		expirationTime := token.ExpirationTime().Unix() - time.Now().Unix()

		ctx.SetCookie("token", token.Token(), int(expirationTime), "/", "127.0.0.1", false, false)

		ctx.JSON(http.StatusOK, response{"ok", ""})
	}
}
