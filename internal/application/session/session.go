package session

import (
	"context"
	"errors"
	"fmt"

	"github.com/vosgaust/voicemod-challenge.git/internal/application/user"
	"github.com/vosgaust/voicemod-challenge.git/internal/domain/session"
	domainuser "github.com/vosgaust/voicemod-challenge.git/internal/domain/user"
)

type SessionService struct {
	userService user.UserService
	auth        session.SessionRepository
}

func NewSessionService(userService user.UserService, auth session.SessionRepository) SessionService {
	return SessionService{userService, auth}
}

func (s SessionService) Authenticate(ctx context.Context, email, password string) (session.Token, error) {
	// TODO: Use new service FindByEmail
	existingUser, err := s.userService.FindByEmail(ctx, email)
	if err != nil {
		fmt.Printf("User not found: %v\n", err)
		return session.Token{}, err
	}
	// Compare password
	if existingUser.Password().String() != domainuser.EncryptPassword(password, existingUser.ID().String()) {
		fmt.Println("Incorrect password")
		return session.Token{}, errors.New("Incorrect password")
	}
	// Generate token
	userSession := session.NewSession(email, password)
	token, err := s.auth.GenerateSession(userSession)
	if err != nil {
		fmt.Printf("Could not generate token: %v\n", err)
		return session.Token{}, err
	}

	fmt.Printf("Generated token: %v\n", token)

	return token, nil
}
