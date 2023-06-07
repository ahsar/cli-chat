package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahsar/cli-chat/app"
)

func exitFunc() {
	os.Exit(0)
}

func main() {
	c := make(chan os.Signal)

	// 监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("exit", s)
				exitFunc()
			default:
				fmt.Println("other", s)
			}
		}
	}()

	fmt.Println("Welcome to use cli-chat")
	app.Run()
}
