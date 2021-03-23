package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vosgaust/voicemod-challenge.git/internal/domain/session"
)

type JWT struct {
	ExpirationTime time.Time
	Key            string
}

func NewJWTAuthentication(timeToExpire int, key string) *JWT {
	expirationTime := time.Now().Add(time.Hour * 24 * time.Duration(timeToExpire))
	return &JWT{expirationTime, key}
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *JWT) GenerateSession(userSession session.Session) (session.Token, error) {
	claims := &Claims{
		Username: userSession.Identity().String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: s.ExpirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.Key))
	//TODO: handle error
	if err != nil {
		return session.Token{}, err
	}

	return session.NewToken(tokenString, s.ExpirationTime), nil
}
