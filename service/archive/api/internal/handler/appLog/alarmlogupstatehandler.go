package appLog

import (
	"net/http"
	"tpmt-zt/common"
	"tpmt-zt/common/responsex"
	"tpmt-zt/service/archive/api/internal/logic/appLog"
	"tpmt-zt/service/archive/api/internal/svc"
	"tpmt-zt/service/archive/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AlarmLogUpStateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AlarmUpStateReqest
		if err := httpx.Parse(r, &req); err != nil {
			responsex.HttpResult(r, w, req, "", common.NewParamError(err.Error()), svcCtx.ArchiveRpc)
			return
		}

		l := appLog.NewAlarmLogUpStateLogic(r.Context(), svcCtx)
		resp, err := l.AlarmLogUpState(&req)
		responsex.HttpResult(r, w, req, resp, err, svcCtx.ArchiveRpc)
	}
}
