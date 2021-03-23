package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Invalid_ID_Error(t *testing.T) {
	userID := "invalid-id"
	// userName := "Ibai"
	// userSurnames := "LLanos"
	// userEmail := "ibai@twitch.com"
	// userPassword := "123456"
	// userCountry := "spain"
	// userPhone := "666666666"
	// userPostalCode := "28008"

	_, err := NewUserID(userID)

	assert.Error(t, err)
}

func Test_Correct_ID(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	// userName := "Ibai"
	// userSurnames := "LLanos"
	// userEmail := "ibai@twitch.com"
	// userPassword := "123456"
	// userCountry := "spain"
	// userPhone := "666666666"
	// userPostalCode := "28008"

	_, err := NewUserID(userID)

	assert.NoError(t, err)
}

func Test_Empty_Name_Error(t *testing.T) {
	_, err := NewUserName("")

	assert.Error(t, err)
}

func Test_Correct_Name(t *testing.T) {
	userName := "Ibai"

	_, err := NewUserName(userName)

	assert.NoError(t, err)
}

func Test_Empy_Surnames_Error(t *testing.T) {
	_, err := NewUserSurnames("")

	assert.Error(t, err)
}

func Test_Correct_Surnames(t *testing.T) {
	surnames := "LLanos"

	_, err := NewUserSurnames(surnames)

	assert.NoError(t, err)
}

func Test_Invalid_Email_Error(t *testing.T) {
	_, err := NewUserEmail("mynewemail.com")

	assert.Error(t, err)
}

func Test_Correct_Email(t *testing.T) {
	email := "ibai@twitch.com"

	_, err := NewUserEmail(email)

	assert.NoError(t, err)
}

func Test_Password_Lenth_Error(t *testing.T) {
	_, err := NewUserPassword("123456")

	assert.Error(t, err)
}

func Test_Correct_Password(t *testing.T) {
	_, err := NewUserPassword("1234567890")

	assert.NoError(t, err)
}

func Test_Create_User_Invalid_ID_Error(t *testing.T) {
	userID := "invalid-id"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibai@twitch.com"
	userPassword := "1234567890"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	_, err := NewUser(
		userID,
		userName,
		userSurnames,
		userEmail,
		userPassword,
		userCountry,
		userPhone,
		userPostalCode)

	assert.Error(t, err)
}

func Test_Create_User_Empty_Name_Error(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := ""
	userSurnames := "LLanos"
	userEmail := "ibai@twitch.com"
	userPassword := "1234567890"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	_, err := NewUser(
		userID,
		userName,
		userSurnames,
		userEmail,
		userPassword,
		userCountry,
		userPhone,
		userPostalCode)

	assert.Error(t, err)
}

func Test_Create_User_Empty_Surnames_Error(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := ""
	userEmail := "ibai@twitch.com"
	userPassword := "1234567890"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	_, err := NewUser(
		userID,
		userName,
		userSurnames,
		userEmail,
		userPassword,
		userCountry,
		userPhone,
		userPostalCode)

	assert.Error(t, err)
}

func Test_Create_User_Invalid_Email_Error(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibaitwitch.com"
	userPassword := "1234567890"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	_, err := NewUser(
		userID,
		userName,
		userSurnames,
		userEmail,
		userPassword,
		userCountry,
		userPhone,
		userPostalCode)

	assert.Error(t, err)
}

func Test_Create_User_Invalid_Pasword_Error(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userPassword := "123456789"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	_, err := NewUser(
		userID,
		userName,
		userSurnames,
		userEmail,
		userPassword,
		userCountry,
		userPhone,
		userPostalCode)

	assert.Error(t, err)
}

func Test_Create_User_Success(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Ibai"
	userSurnames := "LLanos"
	userEmail := "ibait@witch.com"
	userPassword := "1234567890"
	userCountry := "spain"
	userPhone := "666666666"
	userPostalCode := "28008"

	user, err := NewUser(
		userID,
		userName,
		userSurnames,
		userEmail,
		userPassword,
		userCountry,
		userPhone,
		userPostalCode)

	assert.NoError(t, err)
	assert.Equal(t, userID, user.id.value)
	assert.Equal(t, userName, user.name.value)
	assert.Equal(t, userSurnames, user.surnames.value)
	assert.Equal(t, userEmail, user.email.value)
	assert.Equal(t, userPassword, user.hashedPassword.value)
	assert.Equal(t, userPhone, user.phone.value)
	assert.Equal(t, userPostalCode, user.postalCode.value)
}
