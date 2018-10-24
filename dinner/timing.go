package dinner

import (
	"fmt"
	"time"
)

//定时执行任务
func LoopHaveDinner() {
	for {
		now := time.Now()
		// 计算下一个零点
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 14, 10, 0, 0, next.Location())
		fmt.Println("下次执行时间：", next.Format("2006-01-02 15:04:05"))
		t := time.NewTimer(next.Sub(now))
		<-t.C
		LoopHaveDinner()
	}
}
