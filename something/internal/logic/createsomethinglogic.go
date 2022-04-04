package logic

import (
	"context"
	"encoding/json"
	"home-bank/something/internal/svc"
	"home-bank/something/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSomethingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSomethingLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateSomethingLogic {
	return CreateSomethingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSomethingLogic) CreateSomething(req types.CreateRequest) (resp *types.CreateResponse, err error) {
	params := map[string]interface{}{
		"name":         req.Name,
		"category_id":  req.CategoryId,
		"container_id": req.ContainerId,
	}

	params_json, _ := json.Marshal(params)

	resp.Message = string(params_json)
	resp.Code = 22000

	 return resp, nil
}
