package cockroach

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lualfe/casamento/app/responsewriter"
	"github.com/lualfe/casamento/models"
	"github.com/lualfe/casamento/utils"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

// FindUser gets an user information
func (a *DB) FindUser(id uuid.UUID) (*models.User, responsewriter.Response) {
	user := &models.User{}
	if err := a.Instance.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return user, responsewriter.Success()
}

// CreateUser inserts an user into the database
func (a *DB) CreateUser(user *models.User) (*models.User, responsewriter.Response) {
	// validate info
	if len(user.Password) < 6 {
		return nil, responsewriter.Error("password must be at least 6 characters long", http.StatusBadRequest)
	}

	// check if user already exists
	var count int
	if err := a.Instance.Table("users").Where("email = ?", user.Email).Count(&count).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	if count > 0 {
		return nil, responsewriter.Error("user already exists", http.StatusBadRequest)
	}

	// inserts into the database
	user.ID = uuid.New()
	coupleID, _ := a.CreateCouple()
	user.CoupleID = coupleID
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	user.Password = hash
	if err := a.Instance.Create(&user).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return user, responsewriter.Success()
}

// LoginUser checks if the given information checks
func (a *DB) LoginUser(email, password string) (*models.User, responsewriter.Response) {
	user := &models.User{}
	if err := a.Instance.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	if !utils.ComparePassword(password, user.Password) {
		return nil, responsewriter.Error("email or password incorrect", http.StatusUnauthorized)
	}
	return user, responsewriter.Success()
}

// UpdateUser updates an user
func (a *DB) UpdateUser(user *models.User) (*models.User, responsewriter.Response) {
	if err := a.Instance.Save(&user).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return user, responsewriter.Success()
}

// CheckTempPassword verifies if the temporary password matches
func (a *DB) CheckTempPassword(id uuid.UUID, temp string) (bool, responsewriter.Response) {
	user := &models.User{}
	if err := a.Instance.Where("id = ?", id).First(&user).Error; err != nil {
		return false, responsewriter.UnexpectedError(err)
	}
	if user.TemporaryPassword != temp {
		return false, responsewriter.Error("invalid credentials", http.StatusForbidden)
	}
	return true, responsewriter.Success()
}

// DeleteTempPassword Erases temporary password after use
func (a *DB) DeleteTempPassword(id uuid.UUID) responsewriter.Response {
	if err := a.Instance.Table("users").Where("id = ?", id).Update("temporary_password", gorm.Expr("NULL")).Error; err != nil {
		return responsewriter.UnexpectedError(err)
	}
	return responsewriter.Success()
}

// PartnerInvite invites a partner to form a couple id
func (a *DB) PartnerInvite(userID, coupleID uuid.UUID, email, name string) responsewriter.Response {
	user := &models.User{}
	if err := a.Instance.Where("id = ?", userID).First(&user).Error; err != nil {
		return responsewriter.UnexpectedError(err)
	}

	couple := &models.Couple{}
	if err := a.Instance.Select("token").Table("couples").Where("id = ?", user.CoupleID).First(&couple).Error; err != nil {
		return responsewriter.UnexpectedError(err)
	}

	partner := &models.User{}
	if err := a.Instance.Where("email = ?", email).Find(&partner).Error; err != nil {
		if err.Error() != "record not found" {
			return responsewriter.UnexpectedError(err)
		}
	}
	if partner.CoupleID == user.CoupleID {
		return responsewriter.Error("the person you're sending the invite to already is your couple", http.StatusBadRequest)
	}
	link := fmt.Sprintf("%v/user/register?couple=%v&token=%v", viper.GetString("base_url"), user.CoupleID, couple.Token)
	if partner.Email != "" {
		partner.TemporaryPassword = utils.RandStringRunes(20)
		if err := a.Instance.Save(&partner).Error; err != nil {
			return responsewriter.UnexpectedError(err)
		}
		link = fmt.Sprintf("%v/user/update/%v?couple=%v&token=%v&current_couple=%v&temp=%v", viper.GetString("base_url"), partner.ID, user.CoupleID, couple.Token, partner.CoupleID, partner.TemporaryPassword)
	}
	msg := fmt.Sprintf(`Olá %v</br>
	Você recebeu um convite do %v para entrar em nossa plataforma.
	É só acessar o link para entrar!
	<a href="%v">Aqui</a>`, name, user.Name, link)

	m := gomail.NewMessage()
	m.SetHeader("From", "from@example.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", msg)

	d := gomail.Dialer{Host: "localhost", Port: 1025}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return responsewriter.Success()
}
