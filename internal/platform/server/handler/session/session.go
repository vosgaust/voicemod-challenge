package session

import (
	"fmt"
	"net/http"

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
			fmt.Printf("failed creating user: %v", err)
			ctx.JSON(http.StatusUnprocessableEntity, response{"error", "Login payload is incorrect"})
			return
		}

		token, err := sessionService.Authenticate(ctx, login.Email, login.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response{"error", ""})
			return
		}

		ctx.SetCookie("token", token.String(), 5, "/", "127.0.0.1", false, false)

		ctx.JSON(http.StatusOK, response{"ok", ""})
	}
}
