package middleware

import (
	"net/http"
	"strings"

	"github.com/LuanTenorio/learn-api/internal/auth"
	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := getToken(c.Request())

		if err != nil {
			return err
		}

		claims, err := validateToken(token)

		if err != nil {
			return exception.New("Invalid token", http.StatusUnauthorized, err.Error())
		}

		c.Set("claims", claims)

		return next(c)
	}
}

func getToken(r *http.Request) (string, exception.Exception) {
	bearer := r.Header.Get("Authorization")

	if bearer == "" {
		return "", exception.New("Token absent", http.StatusUnauthorized)
	}

	bearerParts := strings.Split(bearer, " ")

	if len(bearerParts) != 2 || bearerParts[0] != "Bearer" {
		return "", exception.New("Invalid token", http.StatusUnauthorized)
	}

	return bearerParts[1], nil
}

func validateToken(tokenString string) (*auth.JwtCustomClaims, exception.Exception) {
	token, err := jwt.ParseWithClaims(tokenString, &auth.JwtCustomClaims{}, keyFunc)

	if err != nil {
		return nil, exception.New("Error when making the token", http.StatusUnauthorized)
	} else if claims, ok := token.Claims.(*auth.JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, exception.New("invalid token", http.StatusUnauthorized)
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, exception.New("invalid signature", http.StatusUnauthorized)
	}

	return auth.SecretKey, nil
}
