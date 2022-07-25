package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issure    string
}

type Claim struct {
	Sum                uint `json:"sum"`
	jwt.StandardClaims `json:"standard_claims"`
}

func NewJwtService() *jwtService {
	return &jwtService{
		secretKey: "secret-key",
		issure:    "book-api",
	}
}

func (s *jwtService) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return " ", err
	}

	return t, nil
}

func (s *jwtService) ValidadeToken(token string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("O token é invalido: ", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}
