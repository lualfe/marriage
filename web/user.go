package web

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/lualfe/casamento/models"
	"github.com/lualfe/casamento/utils"
	"github.com/spf13/viper"
)

// FindUser finds user information
func (w *Web) FindUser(c echo.Context) error {
	return nil
}

// CreateUser creates an user and an session
func (w *Web) CreateUser(c echo.Context) error {
	if c.Get("id") != nil {
		c.String(http.StatusForbidden, "user already logged in")
		return nil
	}
	values, _ := c.FormParams()
	user := &models.User{
		Email:    values.Get("email"),
		Password: values.Get("password"),
	}
	u, r := w.App.Cockroach.CreateUser(user)
	if r.Code == http.StatusOK {
		claims := &utils.JWTClaims{
			ID:    u.ID.String(),
			Email: u.Email,
		}
		claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(viper.GetString("jwt_key")))
		if err != nil {
			return err
		}
		c.Set("id", u.ID)
		c.Set("email", u.Email)
		type finalResponse struct {
			User  *models.User `json:"user"`
			Token string       `json:"access_token"`
		}
		response := &finalResponse{
			User:  u,
			Token: t,
		}
		r.JSON(c, response)
		return nil
	}
	r.JSON(c, "")
	return nil
}

// LoginUser logs the user in
func (w *Web) LoginUser(c echo.Context) error {
	// checks if user is logged
	checkUser := c.Get("user")
	if checkUser != nil {
		return nil
	}

	email := c.FormValue("email")
	password := c.FormValue("password")
	user, r := w.App.Cockroach.LoginUser(email, password)
	if r.Code == http.StatusOK {
		token := jwt.New(jwt.SigningMethodES256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.ID
		claims["email"] = user.Email
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		t, err := token.SignedString([]byte(viper.GetString("jwt_key")))
		if err != nil {
			return err
		}
		type finalResponse struct {
			User  *models.User `json:"user"`
			Token string       `json:"access_token"`
		}
		response := &finalResponse{
			User:  user,
			Token: t,
		}
		r.JSON(c, response)
		return nil
	}
	r.JSON(c, "")
	return nil
}
