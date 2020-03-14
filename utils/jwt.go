package utils

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

// JWTClaims model.
type JWTClaims struct {
	UserID   string `json:"user_id"`
	CoupleID string `json:"couple_id"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// JWTGetter is a helper function that returns the
func JWTGetter(c echo.Context, claims ...string) []interface{} {
	user := c.Get("user").(*jwt.Token)
	claimsList := user.Claims.(jwt.MapClaims)
	var result []interface{}
	for _, claim := range claims {
		result = append(result, claimsList[claim])
	}
	return result
}

// CheckToken is a helper function that checks if a token exists in the session
func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		authSlice := strings.Split(auth, " ")
		if len(authSlice) >= 2 {
			token := authSlice[1]
			claims := &JWTClaims{}
			t, _ := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(viper.GetString("jwt_key")), nil
			})
			if t != nil {
				c.Set("user_id", claims.UserID)
				c.Set("email", claims.Email)
			}
		}
		return next(c)
	}
}
