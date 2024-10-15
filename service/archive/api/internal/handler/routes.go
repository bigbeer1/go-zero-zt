// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package handler

import (
	"net/http"

	appLog "tpmt-zt/service/archive/api/internal/handler/appLog"
	"tpmt-zt/service/archive/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/alarmLog",
				Handler: appLog.AlarmLogHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/alarmLogUpState",
				Handler: appLog.AlarmLogUpStateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/appLog",
				Handler: appLog.AppLogHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/scheduledTasksFailureRecord",
				Handler: appLog.ScheduledTasksFailureRecordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/scheduledTasksLog",
				Handler: appLog.ScheduledTasksLogHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
