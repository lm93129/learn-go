package main

import "time"

func main() {
	var Ball int
	table := make(chan int)
	//用一个通道代表乒乓球台。一个整型变量代表球，然后用两个 goroutine 代表玩家
	//玩家通过增加整型变量的值（点击计数器）模拟击球动作。
	go player(table)
	go player(table)

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(table chan int) {
	for {
		ball := <-table
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
