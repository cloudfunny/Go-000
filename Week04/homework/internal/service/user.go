package service

import (
	"context"
	"strconv"

	"github.com/mohuishou/go-training/Week02/work/pkg/errcode"
	v1 "github.com/mohuishou/go-training/Week04/homework/apis/mohuishou/user/v1"
	"github.com/mohuishou/go-training/Week04/homework/internal/domain"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

// UserService UserService
type UserService struct {
	v1.UserServerServer

	usecase domain.IUserUsecase
}

// NewUserService NewUserService
func NewUserService(usecase domain.IUserUsecase) *UserService {
	return &UserService{usecase: usecase}
}

// GetUserInfo 获取用户信息
func (u *UserService) GetUserInfo(ctx context.Context, req *v1.GetUserInfoRequest) (*v1.GetUserInfoResponse, error) {
	// TODO: 下面这些应该放到中间件中
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.Wrap(errcode.ErrUnKnown, "get metadata err")
	}

	data := md.Get("uid")
	if len(data) != 1 {
		return nil, errors.Wrapf(errcode.ErrUnKnown, "user id lens not 1, metadata: %v", data)
	}

	id, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, errors.Wrapf(errcode.ErrUnKnown, "user id not a num, data: %v", data)
	}

	user, err := u.usecase.GetUserInfo(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &v1.GetUserInfoResponse{
		Username: user.Name,
		City:     user.City,
	}

	return resp, nil
}
