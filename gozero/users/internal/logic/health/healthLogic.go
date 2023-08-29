package health

import (
	"context"

	"users/internal/svc"
	"users/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthLogic) Health(req *types.PingRequest) (resp *types.PingResponse, err error) {

	res := &types.PingResponse{}
	_, err = l.svcCtx.UserModel.Ping()

	if err != nil {
		res.Message = "db down"
	}
	res.Message = "db good"

	return res, nil
}
