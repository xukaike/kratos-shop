package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
	"user/internal/biz"

	"user/api/user/v1"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (u *UserService) CreateUser(ctx context.Context, req *v1.CreateUserInfo) (*v1.UserInfoResponse, error) {
	user, err := u.uc.Create(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		NickName: req.NickName,
	})
	if err != nil {
		return nil, err
	}

	userInfoRsp := v1.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoRsp.Birthday = user.Birthday.Unix()
	}
	return &userInfoRsp, nil
}

func (u *UserService) GetUserList(ctx context.Context, req *v1.PageInfo) (*v1.UserListResponse, error) {
	list, total, err := u.uc.List(ctx, int(req.Pn), int(req.PSize))
	if err != nil {
		return nil, err
	}

	rsp := &v1.UserListResponse{}
	rsp.Total = int32(total)

	for _, user := range list {
		userInfoRsp := UserResponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}

	return rsp, nil
}

func UserResponse(user *biz.User) v1.UserInfoResponse {
	userInfoRsp := v1.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoRsp.Birthday = user.Birthday.Unix()
	}
	return userInfoRsp
}

func (u *UserService) GetUserByMobile(ctx context.Context, req *v1.MobileRequest) (*v1.UserInfoResponse, error) {
	user, err := u.uc.UserByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}

	rsp := UserResponse(user)
	return &rsp, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserInfo) (*emptypb.Empty, error) {
	birthDay := time.Unix(int64(req.Birthday), 0)
	user, err := u.uc.UpdateUser(ctx, &biz.User{
		ID:       req.Id,
		Gender:   req.Gender,
		Birthday: &birthDay,
		NickName: req.NickName,
	})

	if err != nil {
		return nil, err
	}

	if user == false {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (u *UserService) CheckPassword(ctx context.Context, req *v1.PasswordCheckInfo) (*v1.CheckResponse, error) {
	check, err := u.uc.CheckPassword(ctx, req.Password, req.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	return &v1.CheckResponse{Success: check}, nil
}

func (u *UserService) GetUserById(ctx context.Context, req *v1.IdRequest) (*v1.UserInfoResponse, error) {
	user, err := u.uc.UserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	rsp := UserResponse(user)
	return &rsp, nil
}
