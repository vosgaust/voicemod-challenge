package session

type Token struct {
	value string
}

func NewToken(value string) Token {
	return Token{value}
}

func (token Token) String() string {
	return token.value
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

type SessionRepository interface {
	GenerateSession(session Session) (Token, error)
}
