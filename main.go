package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mrNobody95/Gate/core"
	"os"
	"os/signal"
	"syscall"
)

func rec(arr []int) {
	arr[0] = arr[0] * -1
}

func main() {
	//diff:=time.Now().AddDate(0, 0, 1).Sub(time.Now())/time.Hour
	//fmt.Println(diff)
	//return

	exit := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-exit:
			core.Stop()
			done <- true
		}
	}()
	core.StartNewNode("coinex", "")

	<-done
	fmt.Println("exiting")

	return
}
