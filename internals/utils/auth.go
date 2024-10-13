package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"rest/api/internals/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {}

func (a *Auth) HashPassword(password string) (string, error) {
	if len(password) < 8 {
		return "", errors.New("password must be atleast 8 characters long")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (a *Auth) ComparePassword(hashedPassword string, userPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword),)
	if err != nil {
		return err
	}
	return nil
}

func (a *Auth) GenerateToken(payload dto.TokenPayload, secret string) (string, error) {
	unSignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": payload.ID,
		"email": payload.Email,
		"role": payload.Role,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	token, err := unSignedToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *Auth) DecodeToken(token string, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secret), nil
	})
}

func (a *Auth) GenerateRandomCode(codeLength int) (string, error) {
	bytes := make([]byte, codeLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func (a *Auth) GenerateOTP(codeLength int) (string, error) {
	// charset
	const charset = "1234567890"
	// make a slice with size of our codelength
	buffer := make([]byte, codeLength);

	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharLength := len(charset)
	for i := 0; i < codeLength; i++ {
		buffer[i] = charset[int(buffer[i])%otpCharLength];
	}

	return string(buffer), nil
}