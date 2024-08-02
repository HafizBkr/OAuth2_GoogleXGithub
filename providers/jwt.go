package providers

import (
	"context"
	"fmt"
	"os"

	"github.com/go-chi/jwtauth/v5"
)

type JWTProvider struct {
	jwt *jwtauth.JWTAuth
}

func NewJWTProvider() *JWTProvider {
	jwt := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
	return &JWTProvider{jwt: jwt}
}

func (j *JWTProvider) Encode(claims map[string]interface{}) (string, error) {
	token := ""
	_, token, err := j.jwt.Encode(claims)
	if err != nil {
		return token, fmt.Errorf("Error while encoding token: %w", err)
	}
	return token, nil
}

func (j *JWTProvider) Decode(t string) (map[string]interface{}, error) {
	jwt, err := j.jwt.Decode(t)
	if err != nil {
		return nil, fmt.Errorf("Error while decoding token: %w", err)
	}
	claims, err := jwt.AsMap(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Error while converting token data to map: %w", err)
	}
	return claims, nil
}
