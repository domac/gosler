###基于Go实现带参的函数的定时调度

---

具体调用 :

```golang
package main
import (
	"fmt"
	"study/gosler"
)

func task() {
	fmt.Println("测试运行....")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	gosler.Every(1).Second().Do(taskWithParams, 1, "hello")

	gosler.Every(1).Day().At("18:56").Do(task)

	//调度器启动
	gosler.Start()
}
```

并发调用 :

```golang

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

```

