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

type UserGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGetLogic {
	return &UserGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var lg = fmt.Println

func (l *UserGetLogic) UserGet(req *types.UserGetRequest) (resp *types.UserGetResponse, err error) {
	user, err := l.svcCtx.UserModel.FindByName(l.ctx, "yale")

	lg(user)

	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查詢失敗")
	}
	if user == nil {
		return nil, errors.New("查無此人")
	}

	lg(user.Name)
	lg(user.Email)

	return &types.UserGetResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
