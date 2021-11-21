package logic

import (
	"context"
	"errors"
	"strings"

	"github.com/oscto/ky3k"

	"account.oscto.icu/metadata"

	account "account.oscto.icu/proto/handler"
)

func (u *Logic) Register(ctx context.Context, req *account.RegisterRequest, rsp *account.RegisterResponse) error {

	if u.repository.UserEmailExist(req.Email) {
		return errors.New("邮箱已经存在")
	}

	str := strings.Split(req.Email, "@")

	var user metadata.UserCreateBody
	user.Email = req.Email
	user.Salt = ky3k.RandString(8)
	user.Password = ky3k.MarshalMd5(req.Password, user.Salt)
	user.Nickname = str[0]

	data, err := u.repository.UserCreate(user)
	if err != nil {
		return err
	}

	rsp.Username = data.Username
	rsp.Token = data.Email

	return nil
}
