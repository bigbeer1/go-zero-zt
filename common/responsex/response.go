package responsex

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
	"time"
	"tpmt-zt/common/datax"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/service/archive/rpc/archive"
)

// http返回
func HttpResult(r *http.Request, w http.ResponseWriter, req interface{}, resp interface{}, err error, archiveRpc archive.Archive) {
	var isRequest int64
	// 请求判断
	if err != nil {
		// 失败
		isRequest = 0
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		resp = err
	} else {
		// 成功
		isRequest = 1
	}

	// 写入日志
	if archiveRpc != nil {
		// 用户登录信息
		tokenData := jwtx.ParseToken(r.Context())
		// 创建30秒协程写入日志
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		// 协程写入日志
		go WriteLog(ctx, tokenData, isRequest, r, req, resp, archiveRpc)

	}
	httpx.WriteJson(w, http.StatusOK, resp)

}

func WriteLog(ctx context.Context, tokenData *jwtx.TokenData, isRequest int64, r *http.Request, req interface{}, resp interface{}, archiveRpc archive.Archive) {
	var err error
	var reqIp string

	// 先获取Axios代理nginx信息
	reqIp = r.Header.Get("X-Real-Ip")
	if reqIp == "" {
		// 分割获取请求IP  获取9005 nginx 代理IP
		reqIps := strings.Split(r.Header.Get("X-Forwarded-For"), ":")
		if len(reqIps) < 1 {
			reqIp = "0.0.0.0"
		} else {
			// 判断IP 并赋值
			if reqIps[0] == "127.0.0.1" || reqIps[0] == "" {
				reqIp = "0.0.0.0"
			} else {
				reqIp = reqIps[0]
			}
		}

	}

	// 计算运算时间
	var timed int64
	startTime, ok := r.Context().Value("startTime").(int64)
	if ok {
		timed = time.Now().UnixMilli() - startTime
	} else {
		timed = 0
	}

	_, err = archiveRpc.AppLoggerAdd(ctx, &archive.AppLoggerAddReq{
		Uid:              tokenData.Uid,
		CreatedName:      tokenData.NickName,
		Ip:               reqIp,
		InterfaceType:    r.Method,
		InterfaceAddress: r.URL.Path,
		RequestData:      datax.ToString(req),
		IsRequest:        isRequest,
		ResponseData:     datax.ToString(resp),
		Timed:            timed,
	})

	if err != nil {
		logx.WithContext(r.Context()).Errorf("【API-Agent-LOG】 : %+v ", "写入日志超时")
	}

}
