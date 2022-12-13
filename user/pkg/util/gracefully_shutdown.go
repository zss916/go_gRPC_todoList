package util

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// todo signal 系统信号发射和接受器
func GracefullyShutdown(server *http.Server) {
	// 创建系统信号接收器接收关闭信号
	done := make(chan os.Signal, 1)
	/**
	os.Interrupt           -> ctrl+c 的信号
	syscall.SIGINT|SIGTERM -> kill 进程时传递给进程的信号
	*/
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	LogrusObj.Println("closing http server gracefully ...")

	if err := server.Shutdown(context.Background()); err != nil {
		LogrusObj.Fatalln("closing http server gracefully failed: ", err)
	}

}

//后台程序通常启动后就一直运行，如果需要停止通常使用kill 进程号发送信号方式处理，而不是直接kill -9 进程号。所以就需要程序中正确接收信号并处理信号
