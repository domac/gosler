package main
import (
	"fmt"
	"study/gosler"
	"os/signal"
	"os"
)

func task() {
	fmt.Println("测试运行....")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {

	otherScheduler := gosler.NewScheduler()
	otherScheduler.Every(5).Seconds().Do(taskWithParams, 2, "Hello")
	go otherScheduler.Start()

	firstScheduler := gosler.NewScheduler()
	firstScheduler.Every(1).Seconds().Do(taskWithParams, 1, "Good")
	go firstScheduler.Start()

	//终止信号
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill)
	<-done

	firstScheduler.Clear()
	otherScheduler.Clear()

	fmt.Println("调度关闭")
}
