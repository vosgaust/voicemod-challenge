package session

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/vosgaust/voicemod-challenge.git/internal/application/user"
	"github.com/vosgaust/voicemod-challenge.git/internal/domain/session"
	domainuser "github.com/vosgaust/voicemod-challenge.git/internal/domain/user"
)

var ErrIncorrectPassword = errors.New("incorrect password")
var ErrUserNotFound = errors.New("user not found")

type SessionService struct {
	userService user.UserService
	auth        session.SessionRepository
}

func NewSessionService(userService user.UserService, auth session.SessionRepository) SessionService {
	return SessionService{userService, auth}
}

func (s SessionService) Authenticate(ctx context.Context, email, password string) (session.Token, error) {
	existingUser, err := s.userService.FindByEmail(ctx, email)
	if err != nil {
		log.Infof("User not found: %v", err)
		return session.Token{}, ErrUserNotFound
	}
	// Compare password
	if existingUser.Password().String() != domainuser.EncryptPassword(password, existingUser.ID().String()) {
		log.Info("Incorrect password")
		return session.Token{}, ErrIncorrectPassword
	}
	// Generate token
	userSession := session.NewSession(email, password)
	token, err := s.auth.GenerateSession(userSession)
	if err != nil {
		log.Infof("Could not generate token: %v", err)
		return session.Token{}, err
	}

	return token, nil
}
