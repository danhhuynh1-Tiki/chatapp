package jwt

import (
	"chat/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Encode(user domain.User) (string, time.Time, error) {

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &domain.Claims{
		ID:    user.ID.String(),
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(domain.JwtKey)
	return tokenString, expirationTime, err
}
func Decode(token string) (*domain.Claims, error) {
	claims := &domain.Claims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return domain.JwtKey, nil
	})
	if !t.Valid {
		return nil, err
	}
	return claims, nil
}
