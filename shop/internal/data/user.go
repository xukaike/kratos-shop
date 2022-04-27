package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	userService "shop/api/service/user/v1"
	"shop/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (u *userRepo) CreateUser(c context.Context, user *biz.User) (*biz.User, error) {
	createUser, err := u.data.uc.CreateUser(c, &userService.CreateUserInfo{
		NickName: user.NickName,
		Password: user.Password,
		Mobile:   user.Mobile,
	})
	if err != nil {
		return nil, err
	}

	return &biz.User{
		ID:       createUser.Id,
		Mobile:   createUser.Mobile,
		NickName: createUser.NickName,
	}, nil
}

func (u *userRepo) UserByMobile(ctx context.Context, mobile string) (*biz.User, error) {
	user, err := u.data.uc.GetUserByMobile(ctx, &userService.MobileRequest{Mobile: mobile})
	if err != nil {
		return nil, err
	}

	return &biz.User{
		ID:       user.Id,
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Password: user.Password,
	}, nil
}

func (u *userRepo) UserById(ctx context.Context, Id int64) (*biz.User, error) {
	user, err := u.data.uc.GetUserById(ctx, &userService.IdRequest{Id: Id})
	if err != nil {
		return nil, err
	}

	return &biz.User{
		ID:       user.Id,
		Mobile:   user.Mobile,
		NickName: user.NickName,
	}, nil
}

func (u *userRepo) CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error) {
	passRsp, err := u.data.uc.CheckPassword(ctx, &userService.PasswordCheckInfo{
		Password:          password,
		EncryptedPassword: encryptedPassword,
	})
	if err != nil {
		return false, err
	}

	return passRsp.Success, nil
}
