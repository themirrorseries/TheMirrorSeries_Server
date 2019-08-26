package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Second * 2)
	go func() {
		select {
		case <-t.C:
			fmt.Println("执行定时器")
		}
	}()
	a := 10
	//time.Sleep(time.Second * 5)
	if a == 10 {
		fmt.Println("关闭定时器")
		t.Stop()
	}
	time.Sleep(time.Second * 5)
}
