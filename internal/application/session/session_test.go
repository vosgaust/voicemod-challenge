package session

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/vosgaust/voicemod-challenge.git/internal/application/user"
	"github.com/vosgaust/voicemod-challenge.git/internal/domain/session"
	domainuser "github.com/vosgaust/voicemod-challenge.git/internal/domain/user"
	"github.com/vosgaust/voicemod-challenge.git/internal/platform/auth/authmocks"
	"github.com/vosgaust/voicemod-challenge.git/internal/platform/storage/storagemocks"
)

func Test_SessionService_Authenticate_FindUser_Error(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userPassword := "1234567890"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	existingUser, err := domainuser.NewUser(
		id,
		userName,
		userSurnames,
		userEmail,
		userPassword,
		userCountry,
		userPhone,
		userPostalCode)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	sessionRepositoryMock := new(authmocks.SessionRepository)
	userRepositoryMock.On("FindByEmail", mock.Anything, existingUser.Email()).Return(nil, errors.New("database error"))

	userService := user.NewUserService(userRepositoryMock)

	sessionService := NewSessionService(userService, sessionRepositoryMock)

	_, err = sessionService.Authenticate(context.Background(), userEmail, userPassword)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_SessionService_Authenticate_IncorrectPassword_Error(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userHashedPassword := "8502b8a47d4a7ceaa6ff7e6f0ddcd4540ddd4718fe901512ae68ba6d1839aa7d"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	incorrectPassword := "2234567890"

	existingUser, err := domainuser.NewUser(
		id,
		userName,
		userSurnames,
		userEmail,
		userHashedPassword,
		userCountry,
		userPhone,
		userPostalCode)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	sessionRepositoryMock := new(authmocks.SessionRepository)
	userRepositoryMock.On("FindByEmail", mock.Anything, existingUser.Email()).Return(&existingUser, nil)

	userService := user.NewUserService(userRepositoryMock)

	sessionService := NewSessionService(userService, sessionRepositoryMock)

	_, err = sessionService.Authenticate(context.Background(), userEmail, incorrectPassword)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_SessionService_Authenticate_GenerateSessionError(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userHashedPassword := "8502b8a47d4a7ceaa6ff7e6f0ddcd4540ddd4718fe901512ae68ba6d1839aa7d"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	userPassword := "1234567890"

	existingUser, err := domainuser.NewUser(
		id,
		userName,
		userSurnames,
		userEmail,
		userHashedPassword,
		userCountry,
		userPhone,
		userPostalCode)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("FindByEmail", mock.Anything, existingUser.Email()).Return(&existingUser, nil)

	userSession := session.NewSession(userEmail, userPassword)
	sessionRepositoryMock := new(authmocks.SessionRepository)
	sessionRepositoryMock.On("GenerateSession", userSession).Return(session.Token{}, errors.New("generate token error"))

	userService := user.NewUserService(userRepositoryMock)

	sessionService := NewSessionService(userService, sessionRepositoryMock)

	_, err = sessionService.Authenticate(context.Background(), userEmail, userPassword)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_SessionService_Authenticate_Success(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userHashedPassword := "8502b8a47d4a7ceaa6ff7e6f0ddcd4540ddd4718fe901512ae68ba6d1839aa7d"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	userPassword := "1234567890"

	existingUser, err := domainuser.NewUser(
		id,
		userName,
		userSurnames,
		userEmail,
		userHashedPassword,
		userCountry,
		userPhone,
		userPostalCode)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("FindByEmail", mock.Anything, existingUser.Email()).Return(&existingUser, nil)

	userSession := session.NewSession(userEmail, userPassword)
	token := session.NewToken("mytoken", time.Now())

	sessionRepositoryMock := new(authmocks.SessionRepository)
	sessionRepositoryMock.On("GenerateSession", userSession).Return(token, nil)

	userService := user.NewUserService(userRepositoryMock)

	sessionService := NewSessionService(userService, sessionRepositoryMock)

	_, err = sessionService.Authenticate(context.Background(), userEmail, userPassword)

	userRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
