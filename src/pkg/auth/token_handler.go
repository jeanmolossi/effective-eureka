package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type AccessToken struct {
	Token      string
	Expiration time.Time
}

var (
	SECRET = []byte(os.Getenv("SESSION_SECRET"))
)

func NewAccessToken(studentID string) *AccessToken {
	expiration := time.Now().UTC().Local().Add(time.Minute * 10)

	claims := &jwt.StandardClaims{
		ExpiresAt: expiration.Unix(),
		Subject:   studentID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenString, _ := token.SignedString(SECRET)

	return &AccessToken{
		Token:      signedTokenString,
		Expiration: expiration,
	}
}

type RefreshToken struct {
	Token      string
	Expiration time.Time
}

func NewRefreshToken(studentID string) *RefreshToken {
	expiration := time.Now().UTC().Local().Add(time.Hour * 24 * 7)

	claims := &jwt.StandardClaims{
		ExpiresAt: expiration.Unix(),
		Subject:   studentID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenString, _ := token.SignedString(SECRET)

	return &RefreshToken{
		Token:      signedTokenString,
		Expiration: expiration,
	}
}

func DecodeAccessToken(token string) (*AccessToken, error) {
	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	return &AccessToken{
		Token:      token,
		Expiration: time.Unix(claims.ExpiresAt, 0),
	}, nil
}

func DecodeRefreshToken(token string) (*RefreshToken, error) {
	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	return &RefreshToken{
		Token:      token,
		Expiration: time.Unix(claims.ExpiresAt, 0),
	}, nil
}
