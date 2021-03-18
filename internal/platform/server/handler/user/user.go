package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"github.com/vosgaust/voicemod-challenge.git/internal/application/user"
)

type response struct {
	Status string `json:"status"`
	Msg    string `json:"msg,omitempty"`
}

type createUserPayload struct {
	Name       string `json:"name"`
	Surnames   string `json:"surnames"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
	PostalCode string `json:"postal_code"`
}

type updateUserPayload struct {
	Name        string `json:"name"`
	Surnames    string `json:"surnames"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
	PostalCode  string `json:"postal_code"`
}

func CreateHandler(userService user.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newUser := createUserPayload{}

		if err := ctx.BindJSON(&newUser); err != nil {
			log.Errorf("failed creating user: %v", err)
			ctx.JSON(http.StatusUnprocessableEntity, response{"ok", "User payload was incorrect"})
			return
		}

		// TODO: generate this
		userID := "1234567890"

		if err := userService.Create(
			ctx,
			userID,
			newUser.Name,
			newUser.Surnames,
			newUser.Email,
			newUser.Password,
			newUser.Country,
			newUser.Phone,
			newUser.PostalCode); err != nil {
			//TODO: Filter errors and return response based on that error type
			ctx.JSON(http.StatusInternalServerError, response{"error", "Failed to create new user"})
		}

		ctx.JSON(http.StatusOK, response{"ok", ""})
	}
}

func UpdateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("user_id")
		log.Infof("updating user: %s", userID)
		// TODO: call update user service
		ctx.JSON(http.StatusOK, response{"ok", ""})
	}
}

func DeleteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: call update user service
		ctx.JSON(http.StatusOK, response{"ok", ""})
	}
}
