package middleware

import (
	"context"
	"net/http"
	"time"
	"tpmt-zt/common"
	"tpmt-zt/common/authx"
	"tpmt-zt/common/responsex"
	"tpmt-zt/service/archive/rpc/archive"
	"tpmt-zt/service/authentication/authentication"
)

type CheckAuthMiddleware struct {
	authenticationRpc authentication.Authentication
	archiveRpc        archive.Archive
	accessSecret      string
}

func NewCheckAuthMiddleware(authenticationRpc authentication.Authentication, archiveRpc archive.Archive, accessSecret string) *CheckAuthMiddleware {
	return &CheckAuthMiddleware{
		authenticationRpc: authenticationRpc,
		archiveRpc:        archiveRpc,
		accessSecret:      accessSecret,
	}
}

func (m *CheckAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 拿到请求的上下文
		ctx := r.Context()

		ctx = context.WithValue(ctx, "startTime", time.Now().UnixMilli())

		// 更新上下文
		r = r.WithContext(ctx)

		if r.URL.Path != "/login" {
			err := authx.Auth(r, m.authenticationRpc, m.accessSecret)
			if err != nil {
				responsex.HttpResult(r, w, "", "", common.NewCodeError(common.AuthErrorCode, err.Error(), ""), m.archiveRpc)
				return
			}
		}
		next(w, r)
	}
}
