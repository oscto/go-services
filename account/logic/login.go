package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/oscto/ky3k"

	account "account.oscto.icu/proto/handler"
)

func (u *Logic) Login(ctx context.Context, req *account.LoginRequest, rsp *account.LoginResponse) error {

	user, err := u.repository.UserGetByEmail(req.Email)
	if err != nil {
		return errors.New(fmt.Sprintf("用户 %s 没有找到", req.Email))
	}

	password := ky3k.MarshalMd5(req.Password, user.Salt)

	if user.Password != password {
		return errors.New("用户密码不正确")
	}

	rsp.Token = user.Email
	rsp.Username = user.Username

	return nil
}
