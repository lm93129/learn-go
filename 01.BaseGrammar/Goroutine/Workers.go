package Goroutine

//与 fan-in 相反的模式是 fan-out 或者 worker 模式。
//多个 goroutine 可以从单个通道读取，从而在 CPU 内核之间分配大量的工作量 因此是 worker 的名称。
//在 Go 中，此模式易于实现：只需以通道为参数启动多个 goroutine，然后将值发送至该通道 - Go 运行时会自动地进行分配和复用

import (
	"fmt"
	"sync"
	"time"
)

func worker(tasksCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			return
		}
		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Println("processing task", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)

	// 启动多个Worker，初始是阻塞的，因为还没有任务
	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg)
	}

	// 发送任务，然后close通道，表示这就是所有任务了
	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}

	close(tasksCh)
}
