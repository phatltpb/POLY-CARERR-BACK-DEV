package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tuongnguyen1209/poly-career-back/config"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

type JwtConfig struct {
}

type action struct {
	Login         string
	LoginEmployer string
	Verify        string
	ResetPassword string
}
type JwtToken struct {
	jwt.StandardClaims
	Id     int
	Role   int
	Action string
}

type JwtTokenChangeEmail struct {
	jwt.StandardClaims
	Id       int
	OldEmail string
	NewEmail string
}

var JwtAction action = action{
	Login:         "LOGIN",
	LoginEmployer: "LOGIN_EMPLOYER",
	Verify:        "VERIFY",
	ResetPassword: "RESET_PASSWORD",
}

func (r *JwtConfig) Encode(data interface{}, exp time.Duration) (string, error) {
	var (
		config    = config.GetConfig()
		secretKey = config.Jwt.JwtSecretKey
	)

	claims := &jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(exp).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func ExtractToken(bearerToken string) string {

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return bearerToken
}

func (r *JwtConfig) Decode(bearerToken string) (*JwtToken, error) {
	var (
		config    = config.GetConfig()
		secretKey = config.Jwt.JwtSecretKey
	)
	tokenString := ExtractToken(bearerToken)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		jsonStr, err := json.Marshal(claims["data"])
		if err != nil {
			return nil, errors.New(validationMessage.TokenInvalid)
		}
		newData := JwtToken{}
		if err := json.Unmarshal(jsonStr, &newData); err != nil {
			return nil, errors.New(validationMessage.TokenInvalid)
		}

		return &newData, nil
	}
	return nil, errors.New(validationMessage.TokenInvalid)
}
func (r *JwtConfig) DecodeWithInterface(bearerToken string, obj interface{}) (interface{}, error) {
	var (
		config    = config.GetConfig()
		secretKey = config.Jwt.JwtSecretKey
	)
	tokenString := ExtractToken(bearerToken)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		jsonStr, err := json.Marshal(claims["data"])
		if err != nil {
			return nil, errors.New(validationMessage.TokenInvalid)
		}
		if err := json.Unmarshal(jsonStr, obj); err != nil {
			return nil, errors.New(validationMessage.TokenInvalid)
		}

		return obj, nil
	}
	return nil, errors.New(validationMessage.TokenInvalid)
}
