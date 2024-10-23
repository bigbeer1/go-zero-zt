package asynqx

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"time"
)

func NewAsynqClient(c cache.CacheConf) *asynq.Client {
	var redisClusterAddrs []string
	var redisPassword string
	for _, item := range c {
		redisClusterAddrs = append(redisClusterAddrs, item.Host)
		redisPassword = item.Pass
	}

	return asynq.NewClient(asynq.RedisClusterClientOpt{
		Addrs:    redisClusterAddrs,
		Password: redisPassword,
	})
}

func NewAsynqServer(c cache.CacheConf, concurrency int) *asynq.Server {
	var redisClusterAddrs []string
	var redisPassword string
	for _, item := range c {
		redisClusterAddrs = append(redisClusterAddrs, item.Host)
		redisPassword = item.Pass
	}

	// 如果Concurrency为0 默认最大线程数为20
	if concurrency == 0 {
		concurrency = 10
	}
	return asynq.NewServer(
		asynq.RedisClusterClientOpt{
			Addrs:    redisClusterAddrs,
			Password: redisPassword,
		},
		asynq.Config{
			IsFailure: func(err error) bool {
				logx.Errorf("asynq 服务端创建失败 err : %+v \n", err)
				return true
			},
			Concurrency: concurrency, //最大线程数
		},
	)
}

// create scheduler2
func NewScheduler(c cache.CacheConf) *asynq.Scheduler {

	location, _ := time.LoadLocation("Asia/Shanghai")

	var redisClusterAddrs []string
	var redisPassword string
	for _, item := range c {
		redisClusterAddrs = append(redisClusterAddrs, item.Host)
		redisPassword = item.Pass
	}

	return asynq.NewScheduler(
		asynq.RedisClusterClientOpt{
			Addrs:    redisClusterAddrs,
			Password: redisPassword,
		}, &asynq.SchedulerOpts{
			Location: location,
			EnqueueErrorHandler: func(task *asynq.Task, opts []asynq.Option, err error) {
				logx.Errorf(fmt.Sprintf("定时任务 <<<<<<<===>>>>> err : %+v , task : %+v", err, task))
			},
		})
}
