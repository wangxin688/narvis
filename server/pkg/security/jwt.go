// Copyright 2024 wangxin.jeffry@gmail.com
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package security

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wangxin688/narvis/server/config"
	ie "github.com/wangxin688/narvis/server/tools/errors"
	"golang.org/x/crypto/bcrypt"
)

const AuthorizationString string = "Authorization"
const AuthorizationBearer string = "Bearer"

type AccessToken struct {
	TokenType             string    `json:"tokenType"`
	AccessToken           string    `json:"accessToken"`
	ExpiresAt             time.Time `json:"expiresAt"`
	IssuedAt              time.Time `json:"issuedAt"`
	RefreshToken          string    `json:"refreshToken"`
	RefreshTokenExpiresAt time.Time `json:"refreshTokenExpiresAt"`
	RefreshTokenIssuedAt  time.Time `json:"refreshTokenIssuedAt"`
}

type Claims struct {
	UserId    string `json:"userId"`
	Username  string `json:"username"`
	Refreshed bool   `json:"refreshed"`
	jwt.RegisteredClaims
}

func CreateAccessToken(userID, username string, refresh bool, expire time.Duration) (token string, expiresAt time.Time, issuedAt time.Time) {
	now := time.Now()
	expiresAt = now.Add(expire)
	issuedAt = now
	claims := Claims{
		UserId:    userID,
		Username:  username,
		Refreshed: refresh,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
		},
	}
	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ = _token.SignedString([]byte(config.Settings.Jwt.SecretKey))
	return token, expiresAt, issuedAt
}

func GenerateTokenResponse(userID string, username string) *AccessToken {
	expireAccess := time.Duration(config.Settings.Jwt.AccessTokenExpiredMinute) * time.Minute
	accessToken, expiresAt, issuedAt := CreateAccessToken(
		userID, username, false, expireAccess)

	// Create a refresh token
	expireRefresh := time.Duration(config.Settings.Jwt.RefreshTokenExpiredMinute) * time.Minute
	refreshToken, refreshExpiresAt, refreshIssuedAt := CreateAccessToken(
		userID, username, true, expireRefresh)
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
func VerifyAccessToken(tokenString string) (ie.ErrorCode, *Claims) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(_ *jwt.Token) (interface{}, error) {
		return []byte(config.Settings.Jwt.SecretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return ie.CodeAccessTokenInvalid, nil
		}
	}
	claims, ok := token.Claims.(*Claims)
	now := jwt.NewNumericDate(time.Now()).Unix()
	if !ok || !token.Valid {
		if claims.ExpiresAt.Unix() < now {
			return ie.CodeAccessTokenExpired, nil
		}
		if claims.IssuedAt.Unix() > now {
			return ie.CodeAccessTokenInvalid, nil
		}
		if claims.Refreshed {
			return ie.CodeAccessTokenInvalidForRefresh, nil
		}
	}
	return ie.ErrorOk, claims
}

// Verify token and return true if token is valid refresh token
func verifyRefreshToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(_ *jwt.Token) (interface{}, error) {
		return []byte(config.Settings.Jwt.SecretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return false
		}
	}
	claims, ok := token.Claims.(*Claims)
	now := jwt.NewNumericDate(time.Now()).Unix()
	if claims.ExpiresAt.Unix() > now && claims.IssuedAt.Unix() < now {
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
	_, _ = jwt.ParseWithClaims(tokenString, claims, func(_ *jwt.Token) (interface{}, error) {
		return []byte(config.Settings.Jwt.SecretKey), nil
	})
	return GenerateTokenResponse(claims.UserId, claims.Username)
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
