package main

import (
	"fmt"
	"havedinner/dinner"
	"time"
)

func main() {
	p := dinner.Person{}
	LoopHaveDinner(&p)
}

// 定时执行任务
// nohup ./havedinner &

func LoopHaveDinner(p *dinner.Person) {
	for {
		now := time.Now()
		// 计算下一个time
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 14, 10, 0, 0, next.Location())
		fmt.Println("下次执行时间：", next.Format("2006-01-02 15:04:05"))
		t := time.NewTimer(next.Sub(now))
		p.HaveDinner("刘俊榜", "1")
		p.HaveDinner("白栋天", "1")
		p.HaveDinner("薛冬梅", "1")
		p.HaveDinner("张峻峰", "1")
		<-t.C
		LoopHaveDinner(p)
	}
}
