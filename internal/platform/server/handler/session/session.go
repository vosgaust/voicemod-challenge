package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
			fmt.Printf("failed binding user: %v", err)
			ctx.JSON(http.StatusUnprocessableEntity, response{"error", "Login payload is incorrect"})
			return
		}

		token, err := sessionService.Authenticate(ctx, login.Email, login.Password)
		if err != nil {
			fmt.Printf("Failed authenticating user: %v\n", err)
			ctx.JSON(http.StatusInternalServerError, response{"error", ""})
			return
		}

		expirationTime := token.ExpirationTime().Unix() - time.Now().Unix()

		ctx.SetCookie("token", token.Token(), int(expirationTime), "/", "127.0.0.1", false, false)

		ctx.JSON(http.StatusOK, response{"ok", ""})
	}
}
