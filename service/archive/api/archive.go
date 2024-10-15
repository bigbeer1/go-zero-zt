package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	zero_handler "github.com/zeromicro/go-zero/rest/handler"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
	"runtime"
	"time"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/archive/api/internal/config"
	"tpmt-zt/service/archive/api/internal/handler"
	"tpmt-zt/service/archive/api/internal/svc"
)

var configFile = flag.String("f", "etc/archive-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("archiveApi任务使用,cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量

	server := rest.MustNewServer(c.RestConf,
		// token错误拦截
		rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
			httpx.WriteJson(w, http.StatusOK, common.NewCodeError(common.TokenErrorCode, msg.TokenError, err.Error()))
		}),
		// 请求方式错误拦截
		rest.WithNotAllowedHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpx.WriteJson(w, http.StatusOK, common.NewCodeError(common.ReqNotAllCode, msg.ReqNotAllError, nil))
		})),
		// 路由错误拦截
		rest.WithNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpx.WriteJson(w, http.StatusOK, common.NewCodeError(common.ReqRoutesErrorCode, msg.ReqRoutesError, nil))
		})),
	)
	defer server.Stop()

	// 设置日志输出 接口慢时间
	zrpc.SetClientSlowThreshold(time.Second * 2000)
	zero_handler.SetSlowThreshold(time.Second * 2000)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Println(fmt.Sprintf("%s %s:%v", c.Name, c.Host, c.Port))
	server.Start()
}
