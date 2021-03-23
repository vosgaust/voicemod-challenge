package user

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/pbkdf2"
)

var ErrInvalidUserID = errors.New("invalid User ID")

type UserID struct {
	value string
}

func NewUserID(value string) (UserID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return UserID{}, fmt.Errorf("%w: %s", ErrInvalidUserID, value)
	}

	return UserID{
		value: v.String(),
	}, nil
}

func (id UserID) String() string {
	return id.value
}

type UserName struct {
	value string
}

func NewUserName(value string) (UserName, error) {
	return UserName{
		value: value,
	}, nil
}

func (name UserName) String() string {
	return name.value
}

type UserSurnames struct {
	value string
}

func NewUserSurnames(value string) (UserSurnames, error) {
	return UserSurnames{
		value: value,
	}, nil
}

func (surnames UserSurnames) String() string {
	return surnames.value
}

var ErrInvalidUserEmail = errors.New("email format is invalid")

type UserEmail struct {
	value string
}

func NewUserEmail(value string) (UserEmail, error) {
	if !strings.Contains(value, "@") {
		return UserEmail{}, fmt.Errorf("%w: %s", ErrInvalidUserEmail, value)
	}
	return UserEmail{
		value: value,
	}, nil
}

func (email UserEmail) String() string {
	return email.value
}

const minPasswordLen = 10

var ErrInvalidUserPassword = fmt.Errorf("password is not at least %d characters", minPasswordLen)

type UserPassword struct {
	value string
}

func NewUserPassword(value string) (UserPassword, error) {
	if len(value) < minPasswordLen {
		return UserPassword{}, fmt.Errorf("%w", ErrInvalidUserPassword)
	}
	return UserPassword{
		value: value,
	}, nil
}

func (password UserPassword) String() string {
	return password.value
}

type UserCountry struct {
	value string
}

func NewUserCountry(value string) UserCountry {
	return UserCountry{value}
}

func (country UserCountry) String() string {
	return country.value
}

type UserPhone struct {
	value string
}

func NewUserPhone(value string) UserPhone {
	return UserPhone{value}
}

func (phone UserPhone) String() string {
	return phone.value
}

type UserPostalCode struct {
	value string
}

func NewUserPostalCode(value string) UserPostalCode {
	return UserPostalCode{value}
}

func (postalCode UserPostalCode) String() string {
	return postalCode.value
}

type User struct {
	id             UserID
	name           UserName
	surnames       UserSurnames
	email          UserEmail
	hashedPassword UserPassword
	country        UserCountry
	phone          UserPhone
	postalCode     UserPostalCode
}

func NewUser(id, name, surnames, email, password, country, phone, postalCode string) (User, error) {
	idValidated, err := NewUserID(id)
	if err != nil {
		return User{}, err
	}

	nameValidated, err := NewUserName(name)
	if err != nil {
		return User{}, err
	}

	surnamesValidated, err := NewUserSurnames(surnames)
	if err != nil {
		return User{}, err
	}

	emailValidated, err := NewUserEmail(email)
	if err != nil {
		return User{}, err
	}

	passwordValidated, err := NewUserPassword(password)
	if err != nil {
		return User{}, err
	}

	countryValidated := NewUserCountry(country)

	phoneValidated := NewUserPhone(phone)

	postalCodeValidated := NewUserPostalCode(postalCode)

	return User{
		id:             idValidated,
		name:           nameValidated,
		surnames:       surnamesValidated,
		email:          emailValidated,
		hashedPassword: passwordValidated,
		country:        countryValidated,
		phone:          phoneValidated,
		postalCode:     postalCodeValidated,
	}, nil
}

func (user User) GetEncryptedPassword() string {
	return strings.Repeat(user.Password().String(), 2)
}

func EncryptPassword(password, salt string) string {
	dk := pbkdf2.Key([]byte(password), []byte(salt), 4096, 32, sha1.New)
	return fmt.Sprintf("%x", dk)
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Name() UserName {
	return u.name
}

func (u *User) Surnames() UserSurnames {
	return u.surnames
}

func (u *User) Email() UserEmail {
	return u.email
}

func (u *User) Password() UserPassword {
	return u.hashedPassword
}

func (u *User) Country() UserCountry {
	return u.country
}

func (u *User) Phone() UserPhone {
	return u.phone
}

func (u *User) PostalCode() UserPostalCode {
	return u.postalCode
}
