package main

import (
	"fmt"
	"time"
)

func main() {
	n:=time.Time{}

	if n.IsZero() {
		fmt.Println("zero")
	}else {
		fmt.Println("not zero")
	}
}
