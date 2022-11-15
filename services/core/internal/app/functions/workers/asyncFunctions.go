package workers

import (
	"fmt"
	"time"
)

var functionChannels map[string]chan interface{}

type AsyncFunctionsWorker struct {
}

func runner(functionName string) {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			data := <-functionChannels[functionName]
			fmt.Println(data)
		}
	}
}
