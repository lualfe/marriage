package cockroach

import (
	"github.com/google/uuid"
	"github.com/lualfe/casamento/app/responsewriter"
	"github.com/lualfe/casamento/models"
	"github.com/lualfe/casamento/utils"
)

// FindUser gets an user information
func (a *DB) FindUser(id string) (*models.User, responsewriter.Response) {
	user := &models.User{}
	if err := a.Instance.Find(&user).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return user, responsewriter.Success()
}

// CreateUser inserts an user into the database
func (a *DB) CreateUser(user *models.User) (*models.User, responsewriter.Response) {
	// validate info
	if len(user.Password) < 6 {
		return nil, responsewriter.BadRequestError("password must be at least 6 characters long")
	}

	// check if user already exists
	var count int
	if err := a.Instance.Table("users").Where("email = ?", user.Email).Count(&count).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	if count > 0 {
		return nil, responsewriter.BadRequestError("user already exists")
	}

	// inserts into the database
	user.ID = uuid.New()
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
		return nil, responsewriter.BadRequestError("email or password incorrect")
	}
	return user, responsewriter.Success()
}
