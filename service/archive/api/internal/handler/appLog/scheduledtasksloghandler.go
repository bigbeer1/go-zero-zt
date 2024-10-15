package appLog

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"tpmt-zt/common"
	"tpmt-zt/common/responsex"
	"tpmt-zt/service/archive/api/internal/logic/appLog"
	"tpmt-zt/service/archive/api/internal/svc"
	"tpmt-zt/service/archive/api/internal/types"
)

func ScheduledTasksLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ScheduledTasksLogReqest
		if err := httpx.Parse(r, &req); err != nil {
			responsex.HttpResult(r, w, req, "", common.NewParamError(err.Error()), svcCtx.ArchiveRpc)
			return
		}

		l := appLog.NewScheduledTasksLogLogic(r.Context(), svcCtx)
		resp, err := l.ScheduledTasksLog(&req)
		responsex.HttpResult(r, w, req, resp, err, nil)
	}
}
