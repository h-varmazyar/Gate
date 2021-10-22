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
	//cli.Execute()
	//arr:=[]int{1,2,3,4,5,6,7}
	//rec(arr[4:])
	//fmt.Println(arr)
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
