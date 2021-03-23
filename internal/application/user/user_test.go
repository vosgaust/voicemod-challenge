package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/vosgaust/voicemod-challenge.git/internal/domain/user"
	"github.com/vosgaust/voicemod-challenge.git/internal/platform/storage/storagemocks"
)

func Test_UserService_FindUser_RepositoryError(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"

	userID, err := user.NewUserID(id)

	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Find", mock.Anything, userID).Return(nil, errors.New(("database error")))

	userService := NewUserService(userRepositoryMock)
	_, err = userService.Find(context.Background(), id)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_UserService_FindUser_Success(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userPassword := "1234567890"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	userID, err := user.NewUserID(id)
	require.NoError(t, err)

	user, err := user.NewUser(
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
	userRepositoryMock.On("Find", mock.Anything, userID).Return(&user, nil)

	userService := NewUserService(userRepositoryMock)
	foundUser, err := userService.Find(context.Background(), id)

	userRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
	assert.EqualValues(t, user, foundUser)
}

func Test_UserService_FindUserByEmail_RepositoryError(t *testing.T) {
	email := "ibai@twitch.com"

	userEmail, err := user.NewUserEmail(email)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("FindByEmail", mock.Anything, userEmail).Return(nil, errors.New(("database error")))

	userService := NewUserService(userRepositoryMock)
	_, err = userService.FindByEmail(context.Background(), email)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_UserService_FindUserByEmail_IncorrectEmail_Error(t *testing.T) {
	email := "ibaitwitch.com"

	userRepositoryMock := new(storagemocks.UserRepository)

	userService := NewUserService(userRepositoryMock)
	_, err := userService.FindByEmail(context.Background(), email)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_UserService_FindUserByEmail_Success(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userPassword := "1234567890"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	user, err := user.NewUser(
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
	userRepositoryMock.On("FindByEmail", mock.Anything, user.Email()).Return(&user, nil)

	userService := NewUserService(userRepositoryMock)
	foundUser, err := userService.FindByEmail(context.Background(), userEmail)

	userRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
	assert.EqualValues(t, user, foundUser)
}

func Test_CreateUser_RepositoryError(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userPassword := "1234567890"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"
	hashedPassword := "8502b8a47d4a7ceaa6ff7e6f0ddcd4540ddd4718fe901512ae68ba6d1839aa7d"

	user, err := user.NewUser(
		id,
		userName,
		userSurnames,
		userEmail,
		hashedPassword,
		userCountry,
		userPhone,
		userPostalCode)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, user).Return(errors.New("database error"))

	userService := NewUserService(userRepositoryMock)
	err = userService.Create(context.Background(), id, userName, userSurnames, userEmail, userPassword, userCountry, userPhone, userPostalCode)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CreateUser_Success(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userPassword := "1234567890"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"
	hashedPassword := "8502b8a47d4a7ceaa6ff7e6f0ddcd4540ddd4718fe901512ae68ba6d1839aa7d"

	user, err := user.NewUser(
		id,
		userName,
		userSurnames,
		userEmail,
		hashedPassword,
		userCountry,
		userPhone,
		userPostalCode)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, user).Return(nil)

	userService := NewUserService(userRepositoryMock)
	err = userService.Create(context.Background(), id, userName, userSurnames, userEmail, userPassword, userCountry, userPhone, userPostalCode)

	userRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_UpdateUser_RepositoryError(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userPassword := "1234567890"
	newPassword := "0987654321"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"
	hashedPassword := "8502b8a47d4a7ceaa6ff7e6f0ddcd4540ddd4718fe901512ae68ba6d1839aa7d"
	newHashedPassword := "886f6cc76b321c95863ce3a50daae343c3a3c79a29a8b0c1174a151259adc70b"

	existingUser, err := user.NewUser(
		id,
		userName,
		userSurnames,
		userEmail,
		hashedPassword,
		userCountry,
		userPhone,
		userPostalCode)
	require.NoError(t, err)

	updateUser, err := user.NewUser(
		id,
		userName,
		userSurnames,
		userEmail,
		newHashedPassword,
		userCountry,
		userPhone,
		userPostalCode)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Find", mock.Anything, existingUser.ID()).Return(&existingUser, nil)
	userRepositoryMock.On("Update", mock.Anything, updateUser).Return(errors.New("database error"))

	userService := NewUserService(userRepositoryMock)
	err = userService.Update(context.Background(), id, userName, userSurnames, userEmail, userPassword, newPassword, userCountry, userPhone, userPostalCode)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_UpdateUser_FindError(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userPassword := "1234567890"
	newPassword := "0987654321"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"
	hashedPassword := "8502b8a47d4a7ceaa6ff7e6f0ddcd4540ddd4718fe901512ae68ba6d1839aa7d"

	existingUser, err := user.NewUser(
		id,
		userName,
		userSurnames,
		userEmail,
		hashedPassword,
		userCountry,
		userPhone,
		userPostalCode)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Find", mock.Anything, existingUser.ID()).Return(nil, errors.New("database error"))

	userService := NewUserService(userRepositoryMock)
	err = userService.Update(context.Background(), id, userName, userSurnames, userEmail, userPassword, newPassword, userCountry, userPhone, userPostalCode)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_UpdateUser_IncorrectPassword_Error(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	newPassword := "0987654321"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"
	hashedPassword := "8502b8a47d4a7ceaa6ff7e6f0ddcd4540ddd4718fe901512ae68ba6d1839aa7d"
	incorrectPassword := "2234567890"

	existingUser, err := user.NewUser(
		id,
		userName,
		userSurnames,
		userEmail,
		hashedPassword,
		userCountry,
		userPhone,
		userPostalCode)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Find", mock.Anything, existingUser.ID()).Return(&existingUser, nil)

	userService := NewUserService(userRepositoryMock)
	err = userService.Update(context.Background(), id, userName, userSurnames, userEmail, incorrectPassword, newPassword, userCountry, userPhone, userPostalCode)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_UpdateUser_IncorrectUser_Error(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	incorrectUserEmail := "ibaitwitch.com"
	userPassword := "2234567890"
	newPassword := "0987654321"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	userRepositoryMock := new(storagemocks.UserRepository)

	userService := NewUserService(userRepositoryMock)
	err := userService.Update(context.Background(), id, userName, userSurnames, incorrectUserEmail, userPassword, newPassword, userCountry, userPhone, userPostalCode)

	assert.Error(t, err)
}

func Test_UpdateUser_Success(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userPassword := "1234567890"
	newPassword := "0987654321"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"
	hashedPassword := "8502b8a47d4a7ceaa6ff7e6f0ddcd4540ddd4718fe901512ae68ba6d1839aa7d"
	newHashedPassword := "886f6cc76b321c95863ce3a50daae343c3a3c79a29a8b0c1174a151259adc70b"

	existingUser, err := user.NewUser(
		id,
		userName,
		userSurnames,
		userEmail,
		hashedPassword,
		userCountry,
		userPhone,
		userPostalCode)
	require.NoError(t, err)

	updateUser, err := user.NewUser(
		id,
		userName,
		userSurnames,
		userEmail,
		newHashedPassword,
		userCountry,
		userPhone,
		userPostalCode)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Find", mock.Anything, existingUser.ID()).Return(&existingUser, nil)
	userRepositoryMock.On("Update", mock.Anything, updateUser).Return(nil)

	userService := NewUserService(userRepositoryMock)
	err = userService.Update(context.Background(), id, userName, userSurnames, userEmail, userPassword, newPassword, userCountry, userPhone, userPostalCode)

	userRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_DeleteUser_RepositoryError(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"

	userID, err := user.NewUserID(id)

	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Delete", mock.Anything, userID).Return(errors.New(("database error")))

	userService := NewUserService(userRepositoryMock)
	err = userService.Delete(context.Background(), id)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_DeleteUser_IncorrectID_Error(t *testing.T) {
	incorrectID := "37a0f027-15e6-47cc-a5d264183281087e"

	userRepositoryMock := new(storagemocks.UserRepository)

	userService := NewUserService(userRepositoryMock)
	err := userService.Delete(context.Background(), incorrectID)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_DeleteUser_Success(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"

	userID, err := user.NewUserID(id)

	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Delete", mock.Anything, userID).Return(nil)

	userService := NewUserService(userRepositoryMock)
	err = userService.Delete(context.Background(), id)

	userRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
