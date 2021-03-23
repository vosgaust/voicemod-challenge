package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/vosgaust/voicemod-challenge.git/internal/application/user"
	domainuser "github.com/vosgaust/voicemod-challenge.git/internal/domain/user"
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
			ctx.JSON(http.StatusUnprocessableEntity, response{"error", "User payload was incorrect"})
			return
		}

		log.Debug("Requesting to create new user")
		userID := uuid.NewV1()

		if err := userService.Create(
			ctx,
			userID.String(),
			newUser.Name,
			newUser.Surnames,
			newUser.Email,
			newUser.Password,
			newUser.Country,
			newUser.Phone,
			newUser.PostalCode); err != nil {
			log.Error("Failed to create new user")
			switch {
			case errors.Is(err, domainuser.ErrInvalidUserEmail):
				ctx.JSON(http.StatusBadRequest, response{"error", "email format is invalid"})
			case errors.Is(err, domainuser.ErrInvalidUserPassword):
				ctx.JSON(http.StatusBadRequest, response{"error", "password format is invalid"})

			default:
				ctx.JSON(http.StatusInternalServerError, response{"error", ""})
				return
			}
		}

		ctx.JSON(http.StatusOK, response{"ok", ""})
	}
}

func UpdateHandler(userService user.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("user_id")

		updatedUser := updateUserPayload{}

		log.Infof("updating user: %s", userID)
		if err := ctx.BindJSON(&updatedUser); err != nil {
			log.Errorf("failed creating user: %v", err)
			ctx.JSON(http.StatusUnprocessableEntity, response{"error", "User payload was incorrect"})
			return
		}

		err := userService.Update(ctx,
			userID,
			updatedUser.Name,
			updatedUser.Surnames,
			updatedUser.Email,
			updatedUser.Password,
			updatedUser.NewPassword,
			updatedUser.Country,
			updatedUser.Phone,
			updatedUser.PostalCode)

		if err != nil {
			switch {
			case errors.Is(err, domainuser.ErrInvalidUserEmail):
				ctx.JSON(http.StatusBadRequest, response{"error", "email format is invalid"})
				return
			case errors.Is(err, domainuser.ErrInvalidUserPassword):
				ctx.JSON(http.StatusBadRequest, response{"error", "password format is invalid"})
				return
			case errors.Is(err, user.ErrIncorrectPassword):
				ctx.JSON(http.StatusBadRequest, response{"error", "invalid password"})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, response{"error", ""})
				return
			}
		}
		ctx.JSON(http.StatusOK, response{"ok", ""})
	}
}

func DeleteHandler(userService user.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("user_id")

		err := userService.Delete(ctx, userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response{"error", "Failed to update user"})
			return
		}

		ctx.JSON(http.StatusOK, response{"ok", ""})
	}
}
