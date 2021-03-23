package session

import "time"

type Token struct {
	token          string
	expirationTime time.Time
}

func NewToken(value string, expirationTime time.Time) Token {
	return Token{value, expirationTime}
}

func (token Token) Token() string {
	return token.token
}

func (token Token) ExpirationTime() time.Time {
	return token.expirationTime
}

type SessionIdentity struct {
	value string
}

func NewSessionIdentity(value string) SessionIdentity {
	return SessionIdentity{value}
}

func (identity SessionIdentity) String() string {
	return identity.value
}

type SessionPassword struct {
	value string
}

func NewSessionPassword(value string) SessionPassword {
	return SessionPassword{value}
}

func (password *SessionPassword) String() string {
	return password.value
}

type Session struct {
	identity SessionIdentity
	password SessionPassword
}

func NewSession(identity, password string) Session {
	validatedIdentity := NewSessionIdentity(identity)
	validatedPassword := NewSessionPassword(password)

	return Session{validatedIdentity, validatedPassword}
}

func (s *Session) Identity() SessionIdentity {
	return s.identity
}

func (s *Session) Password() SessionPassword {
	return s.password
}

//go:generate mockery --case underscore --output ../../platform/auth/authmocks --outpkg authmocks --name SessionRepository
type SessionRepository interface {
	GenerateSession(session Session) (Token, error)
}
