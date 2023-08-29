package user

import (
	"context"
	"errors"
	"fmt"

	"users/internal/model"
	"users/internal/svc"
	"users/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersLogic {
	return &UsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsersLogic) Users(req *types.Request) (resp *types.Response, err error) {

	fmt.Println("hihi")

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, 1)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("failed")
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return &types.Response{
		Message: "hi" + user.Name,
	}, nil
}
