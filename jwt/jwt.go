package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type IJwt interface {
	GenerateToken(issuer, payload string) (string, int64, error)
	GenerateRefreshToken(issuer, payload string) (string, int64, error)
	GetTokenSecret() string
}

type JwtCredential struct {
	TokenSecret         string
	ExpiredToken        int
	RefreshTokenSecret  string
	ExpiredRefreshToken int
}

func (cred JwtCredential) GetTokenSecret() string {
	return cred.TokenSecret
}

func NewJwt(tokenSecret,refreshTokenSecret string, expiredToken,expiredRefreshToken int) IJwt{
	return JwtCredential{
		TokenSecret:         tokenSecret,
		ExpiredToken:        expiredToken,
		RefreshTokenSecret:  refreshTokenSecret,
		ExpiredRefreshToken: expiredRefreshToken,
	}
}

type CustomClaims struct {
	jwt.StandardClaims
	Payload string `json:"payload"`
}

func (cred JwtCredential) GenerateToken(issuer, payload string) (string, int64, error) {
	expirationTime := time.Now().Add(time.Duration(cred.ExpiredToken) * time.Hour).Unix()

	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    issuer,
		},
		Payload: payload,
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(cred.TokenSecret))

	return token, expirationTime, err
}

func (cred JwtCredential) GenerateRefreshToken(issuer, payload string) (string, int64, error) {
	expirationTime := time.Now().Add(time.Duration(cred.ExpiredRefreshToken) * time.Hour).Unix()

	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    issuer,
		},
		Payload: payload,
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(cred.RefreshTokenSecret))

	return token, expirationTime, err
}
