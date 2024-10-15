package sysUser

import (
	"net/http"
	"tpmt-zt/common/responsex"

	"tpmt-zt/service/tpmt/api/internal/logic/sysUser"
	"tpmt-zt/service/tpmt/api/internal/svc"
)

func SysUserLoginInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sysUser.NewSysUserLoginInfoLogic(r.Context(), svcCtx)
		resp, err := l.SysUserLoginInfo()
		responsex.HttpResult(r, w, "", resp, err, nil)
	}
}
