package web

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/lualfe/casamento/app/responsewriter"
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
	// checks if user is logged in
	if c.Get("user_id") != nil {
		response := responsewriter.Redirect("user already logged in", http.StatusMovedPermanently)
		response.JSON(c, "")
		return nil
	}

	values, _ := c.FormParams()
	user := &models.User{
		Email:    values.Get("email"),
		Password: values.Get("password"),
		Name:     values.Get("name"),
	}
	cID := c.QueryParam("couple")
	token := c.QueryParam("token")
	if cID != "" {
		coupleID, _ := uuid.Parse(cID)
		ok, _ := w.App.Cockroach.CheckCoupleToken(token, coupleID)
		if ok {
			user.CoupleID = coupleID
		}
	}
	u, r := w.App.Cockroach.CreateUser(user)
	if r.Code == http.StatusOK {
		claims := &utils.JWTClaims{
			UserID:   u.ID.String(),
			CoupleID: u.CoupleID.String(),
			Email:    u.Email,
		}
		claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(viper.GetString("jwt_key")))
		if err != nil {
			return err
		}
		c.Set("user_id", u.ID)
		c.Set("couple_id", user.CoupleID)
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
	if c.Get("user_id") != nil {
		response := responsewriter.Redirect("user already logged in", http.StatusMovedPermanently)
		response.JSON(c, "")
		return nil
	}

	email := c.FormValue("email")
	password := c.FormValue("password")
	user, r := w.App.Cockroach.LoginUser(email, password)
	if r.Code == http.StatusOK {
		claims := &utils.JWTClaims{
			UserID:   user.ID.String(),
			Email:    user.Email,
			CoupleID: user.CoupleID.String(),
		}
		claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(viper.GetString("jwt_key")))
		if err != nil {
			return err
		}
		c.Set("user_id", user.ID)
		c.Set("couple_id", user.CoupleID)
		c.Set("email", user.Email)
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

// UpdateUser updates user info
func (w *Web) UpdateUser(c echo.Context) error {
	p := c.QueryParams()
	cID := p.Get("couple")
	token := p.Get("token")
	currentCouple := p.Get("current_couple")
	temporaryPassowrd := p.Get("temp")
	uID := c.Param("id")
	if cID != "" {
		coupleID, _ := uuid.Parse(cID)
		tokenOk, response := w.App.Cockroach.CheckCoupleToken(token, coupleID)
		if response.Code != 200 {
			response.JSON(c, "")
			return nil
		}
		userID, _ := uuid.Parse(uID)
		tempPasswordOk, response := w.App.Cockroach.CheckTempPassword(userID, temporaryPassowrd)
		if response.Code != 200 {
			response.JSON(c, "")
			return nil
		}
		if tokenOk && tempPasswordOk {
			user, response := w.App.Cockroach.FindUser(userID)
			if response.Code != 200 {
				response.JSON(c, "")
				return nil
			}
			if user.CoupleID == coupleID {
				response := responsewriter.Error("this couple is already set in your account", http.StatusBadRequest)
				response.JSON(c, "")
				return nil
			}
			user.CoupleID = coupleID
			user, response = w.App.Cockroach.UpdateUser(user)
			response.JSON(c, user)
			if response.Code == 200 {
				claims := &utils.JWTClaims{
					UserID:   user.ID.String(),
					Email:    user.Email,
					CoupleID: user.CoupleID.String(),
				}
				claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				_, err := token.SignedString([]byte(viper.GetString("jwt_key")))
				if err != nil {
					return err
				}
				c.Set("user_id", user.ID)
				c.Set("couple_id", user.CoupleID)
				c.Set("email", user.Email)
				cCouple, _ := uuid.Parse(currentCouple)
				w.App.Cockroach.DeleteCouple(cCouple)
				w.App.Cockroach.DeleteTempPassword(user.ID)
			}
		}
	}
	return nil
}

// PartnerInvite invites a partner to form a couple id
func (w *Web) PartnerInvite(c echo.Context) error {
	jwt := utils.JWTGetter(c, "user_id", "couple_id")
	uID := jwt[0].(string)
	cID := jwt[1].(string)
	partnerEmail := c.FormValue("email")
	partnerName := c.FormValue("name")
	userID, _ := uuid.Parse(uID)
	coupleID, _ := uuid.Parse(cID)
	response := w.App.Cockroach.PartnerInvite(userID, coupleID, partnerEmail, partnerName)
	response.JSON(c, "")
	return nil
}
