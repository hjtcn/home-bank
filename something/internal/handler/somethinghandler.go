package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"home-bank/something/internal/logic"
	"home-bank/something/internal/svc"
	"home-bank/something/internal/types"
)

func SomethingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSomethingLogic(r.Context(), svcCtx)
		resp, err := l.Something(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
