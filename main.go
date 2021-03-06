package main

import (
	"bluebell/common/redis"
	"bluebell/config"
	"bluebell/pkg/snowflake"
	"bluebell/pkg/translator"
	"bluebell/pkg/viper"
	"bluebell/pkg/zaplogger"
	"bluebell/router"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var confPath string

func init() {
	flag.StringVar(&confPath, "config", "./config/config.yaml", "default config path.")
}

// @title bluebell
// @version 1.0
// @description 描述啥
// @termsOfService http://baidu.com/
// @contact.name cwaves
// @contact.url http://www.swagger.io/support
// @contact.email xianchao.chen@foxmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8082
// @BasePath v1
func main() {
	flag.Parse()
	if err := viper.Init(confPath); err != nil {
		fmt.Printf("Init config failed, err:%v\n", err)
		return
	}

	if config.GlobalConfig.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := zaplogger.Init(config.GlobalConfig.LogConfig); err != nil {
		fmt.Printf("Init zaplogger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()

	if err := redis.Init(config.GlobalConfig.RedisConfig); err != nil {
		fmt.Printf("Init Redis failed, err:%v\n", err)
		return
	}

	if err := snowflake.Init(config.GlobalConfig.SnowflakeConfig.StartTime,
		config.GlobalConfig.SnowflakeConfig.MachineID); err != nil {
		fmt.Printf("Init snowflake  failed, err:%v\n", err)
		return
	}

	if err := translator.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	r := gin.New()
	r.Use(zaplogger.GinLogger(), zaplogger.GinRecovery(true))
	router.Register(r)

	pprof.Register(r)

	// 启动服务 优雅关机
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GlobalConfig.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown...", zap.Error(err))
	}

	zap.L().Info("Server exciting ...")
}
