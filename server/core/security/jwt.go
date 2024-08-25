package security

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/tools/errors"
	"golang.org/x/crypto/bcrypt"
)

const AuthorizationString string = "Authorization"
const AuthorizationBearer string = "Bearer"

type AccessToken struct {
	TokenType             string `json:"token_type"`
	AccessToken           string `json:"access_token"`
	ExpiresAt             int64  `json:"expires_at"`
	IssuedAt              int64  `json:"issued_at"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
	RefreshTokenIssuedAt  int64  `json:"refresh_token_issued_at"`
}

type Claims struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Refreshed bool   `json:"refreshed"`
	jwt.StandardClaims
}

func CreateAccessToken(userID string, username string, refresh bool, expire time.Duration) (token string, expiresAt int64, issuedAt int64) {
	now := time.Now()
	expiresAt = now.Add(expire).Unix()
	issuedAt = now.Unix()
	claims := Claims{
		UserID:    userID,
		Username:  username,
		Refreshed: refresh,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
	}
	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ = _token.SignedString([]byte(core.Settings.Jwt.SecretKey))
	return token, expiresAt, issuedAt
}

func GenerateTokenResponse(userID string, username string) *AccessToken {
	accessToken, expiresAt, issuedAt := CreateAccessToken(
		userID, username, false, time.Duration(core.Settings.Jwt.AccessTokenExpiredMinute))

	// Create a refresh token
	refreshToken, refreshExpiresAt, refreshIssuedAt := CreateAccessToken(
		userID, username, true, time.Duration(core.Settings.Jwt.RefreshTokenExpiredMinute))
	return &AccessToken{
		TokenType:             AuthorizationBearer,
		AccessToken:           accessToken,
		ExpiresAt:             expiresAt,
		IssuedAt:              issuedAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshExpiresAt,
		RefreshTokenIssuedAt:  refreshIssuedAt,
	}
}

// verify token and return true if token is valid access token
func VerifyAccessToken(tokenString string) (errors.ErrorCode, *Claims) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(core.Settings.Jwt.SecretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return errors.CodeAccessTokenInvalid, nil
		}
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return errors.CodeAccessTokenInvalid, nil
	}
	if claims.ExpiresAt > time.Now().Unix() {
		return errors.CodeAccessTokenExpired, nil
	}
	if claims.IssuedAt < time.Now().Unix() {
		return errors.CodeAccessTokenInvalid, nil
	}
	if claims.Refreshed {
		return errors.CodeAccessTokenInvalidForRefresh, nil
	}
	return errors.ErrorOk, claims
}

// Verify token and return true if token is valid refresh token
func verifyRefreshToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(core.Settings.Jwt.SecretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false
		}
	}
	claims, ok := token.Claims.(*Claims)
	if claims.ExpiresAt > time.Now().Unix() && claims.IssuedAt < time.Now().Unix() {
		return false
	}
	if !ok || !token.Valid {
		return false
	}
	return claims.Refreshed
}

// Generate token from refresh token
func GenerateRefreshTokenResponse(tokenString string) *AccessToken {
	if !verifyRefreshToken(tokenString) {
		return &AccessToken{}
	}
	claims := &Claims{}
	_, _ = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(core.Settings.Jwt.SecretKey), nil
	})
	return GenerateTokenResponse(claims.UserID, claims.Username)
}

// Hash password
func GetPasswordHash(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// Verify password
func VerifyPasswordHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// Generate random token
func RandomTokenString(n int) (token string) {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	token = hex.EncodeToString(b)
	return
}
