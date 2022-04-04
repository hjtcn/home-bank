package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"home-bank/something/internal/logic"
	"home-bank/something/internal/svc"
	"home-bank/something/internal/types"
)

func CreateSomethingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateSomethingLogic(r.Context(), svcCtx)
		resp, err := l.CreateSomething(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
