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

func AlarmLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AlarmLogReqest
		if err := httpx.Parse(r, &req); err != nil {
			responsex.HttpResult(r, w, req, "", common.NewParamError(err.Error()), svcCtx.ArchiveRpc)
			return
		}

		l := appLog.NewAlarmLogLogic(r.Context(), svcCtx)
		resp, err := l.AlarmLog(&req)
		responsex.HttpResult(r, w, req, resp, err, nil)
	}
}
