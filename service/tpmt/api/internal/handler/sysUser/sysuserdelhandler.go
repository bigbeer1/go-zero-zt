package sysUser

import (
	"net/http"
	"tpmt-zt/common/responsex"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tpmt-zt/service/tpmt/api/internal/logic/sysUser"
	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"tpmt-zt/common"
)

func SysUserDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysUserDelRequest
		if err := httpx.Parse(r, &req); err != nil {
			responsex.HttpResult(r, w, req, "", common.NewParamError(err.Error()), svcCtx.ArchiveRpc)
			return
		}

		l := sysUser.NewSysUserDelLogic(r.Context(), svcCtx)
		resp, err := l.SysUserDel(&req)
		responsex.HttpResult(r, w, req, resp, err, svcCtx.ArchiveRpc)
	}
}
