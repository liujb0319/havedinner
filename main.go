package main

import (
	"fmt"
	"havedinner/dinner"
	"math/rand"
	"strings"

	//"regexp"
	"time"
)

func main() {
	p := dinner.Person{}
	LoopHaveDinner(&p)

}
func Between(str, starting, ending string) string {
	s := strings.Index(str, starting)
	if s < 0 {
		return ""
	}
	s += len(starting)
	e := strings.Index(str[s:], ending)
	if e < 0 {
		return ""
	}
	return str[s : s+e]
}

// create password
func makepw() {

	pwl := 6 //length
	chars := "abcdefghijkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ23456789"
	clen := float64(len(chars))
	res := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < pwl; i++ {
		rfi := int(clen * rand.Float64())
		res += fmt.Sprintf("%c", chars[rfi])
	}
	fmt.Println(res)
}

// 定时执行任务
func LoopHaveDinner(p *dinner.Person) {
	for {
		now := time.Now()
		// 计算下一个time
		next := now.Add(time.Hour * 24)
		//next = time.Date(next.Year(), next.Month(), next.Day(), 14, 10, 0, 0, next.Location())
		next = time.Date(next.Year(), next.Month(), next.Day(), 15, 00, 0, 0, next.Location())
		fmt.Println("下次执行时间：", next.Format("2006-01-02 15:04:05"))
		t := time.NewTimer(next.Sub(now))
		p.HaveDinner("刘俊榜", "1", "1")
		p.HaveDinner("白栋天", "1", "6")
		<-t.C
		LoopHaveDinner(p)
	}
}
