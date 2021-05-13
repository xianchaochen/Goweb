package main

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/routes"
	"web_app/settings"
)

var confPath string

func init()  {
	//./web_app --config=./conf/config.yaml
	flag.StringVar(&confPath, "config", "./conf/config.yaml", "default config path.")
}

func main() {
	flag.Parse()
	if err := settings.Init(confPath); err != nil {
		fmt.Printf("Init config failed, err:%v\n", err)
		return
	}

	// 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("Init logger failed, err:%v\n", err)
		return
	}
	zap.L().Debug("init logger success\n")

	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("Init mysql failed, err:%v\n", err)
		return
	}

	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("Init redis  failed, err:%v\n", err)
		return
	}
	// 缓存区的日志刷出来
	defer redis.Close()
	defer zap.L().Sync()

	r := routes.SetUp()

	// 启动服务 优雅关机
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
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
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)  // 此处不会阻塞
	<-quit  // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown...", zap.Error(err))
	}

	zap.L().Info("Server exciting ...")

	return
}


