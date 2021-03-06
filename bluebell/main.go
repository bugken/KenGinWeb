package main

import (
	"NetClassGinWeb/bluebell/controller"
	"NetClassGinWeb/bluebell/dao/mysql"
	"NetClassGinWeb/bluebell/dao/redis"
	"NetClassGinWeb/bluebell/logger"
	"NetClassGinWeb/bluebell/routes"
	"NetClassGinWeb/bluebell/settings"
	"NetClassGinWeb/bluebell/thirdparty/snowflake"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// @title bluebell项目接口文档
// @version 1.0
// @description Go web开发进阶项目实战课程bluebell

// @contact.name liwenzhou
// @contact.url http://www.liwenzhou.com

// @host 127.0.0.1:8084
// @BasePath /api/v1
func main() {
	//Go Web开发通用的脚手架模板
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("config init failed, error: %s\n", err)
		return
	}

	// 2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("logger init failed, error: %s\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success.")

	// 3.初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		zap.L().Error("mysql init failed, error: %s\n", zap.Error(err))
		return
	}
	defer mysql.Close()

	// 4.初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		zap.L().Error("redis init failed, error: %s\n", zap.Error(err))
		return
	}
	defer redis.Close()

	// 5.初始化snowflake
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("start time:%v\n", settings.Conf.StartTime)
		zap.L().Error("snowflake init failed, error: %s\n", zap.Error(err))
		return
	}

	// 初始化Gin框架内的校验器的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		zap.L().Error("validator trans init failed, error: %s\n", zap.Error(err))
		return
	}

	// 6.注册路由
	r := routes.Setup(settings.Conf.Mode)

	// 7.启动服务(优雅关机/优雅重启)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen error: %s\n", zap.Error(err))
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
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

	return
}
