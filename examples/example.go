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
