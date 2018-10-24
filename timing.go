package main

import (
	"fmt"
	"time"
)

//定时器
func timing1(frequency int) string {
	ticker := time.NewTicker(time.Duration(int64(time.Second) * int64(frequency)))
	for {
		time := <-ticker.C
		fmt.Println("定时器====>", time.String())
	}
}

//在一定时间内接收不到a的数据则超时
func getAchan(timeout time.Duration, a chan string) {
	var after <-chan time.Time
loop:
	after = time.After(timeout)
	for {
		fmt.Println("等待a中的数据，10秒后没有数据则超时")
		select {
		case x := <-a:
			fmt.Println(x)
			goto loop
		case <-after:
			fmt.Println("timeout.")
			return
		}
	}
}

func timing() {
	//定时器，10秒钟执行一次
	ticker := time.NewTicker(10 * time.Second)
	for {
		time := <-ticker.C
		fmt.Println("定时器====>", time.String())
	}
}
