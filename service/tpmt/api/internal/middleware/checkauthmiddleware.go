package middleware

import (
	"context"
	"net/http"
	"time"
)

type CheckAuthMiddleware struct {
}

func NewCheckAuthMiddleware() *CheckAuthMiddleware {
	return &CheckAuthMiddleware{}
}

func (m *CheckAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 拿到请求的上下文
		ctx := r.Context()

		ctx = context.WithValue(ctx, "startTime", time.Now().UnixMilli())

		// 更新上下文
		r = r.WithContext(ctx)

		next(w, r)
	}
}
