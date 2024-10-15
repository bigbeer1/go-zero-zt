package handler

import (
	"net/http"
	"tpmt-zt/common/responsex"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tpmt-zt/service/tpmt/api/internal/logic"
	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"tpmt-zt/common"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			responsex.HttpResult(r, w, req, "", common.NewParamError(err.Error()), svcCtx.ArchiveRpc)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		req.Password = ""
		responsex.HttpResult(r, w, req, resp, err, svcCtx.ArchiveRpc)
	}
}
