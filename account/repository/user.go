package repository

import (
	"errors"
	"fmt"

	"account.oscto.icu/metadata"

	"account.oscto.icu/model"
)

func (r *Repository) UserCreate(data metadata.UserCreateBody) (*model.Users, error) {

	var user = &model.Users{}

	user.Email = data.Email
	user.Salt = data.Salt
	user.Username = data.Username
	user.Password = data.Password
	user.Nickname = data.Nickname

	err := r.db.Create(user).Error

	return user, err
}

func (r *Repository) UserGetByEmail(email string) (model.Users, error) {

	var user model.Users
	r.db.Where("email = ?", email).First(&user)
	var err error
	if user.ID == 0 {
		err = errors.New(fmt.Sprintf("用户 %s 没有找到", email))
	}

	return user, err
}

func (r *Repository) UserEmailExist(email string) bool {

	var total int64
	r.db.Model(&model.Users{}).Where("email = ?", email).Count(&total)

	var has bool
	if total > 0 {
		has = true
	}

	return has
}
