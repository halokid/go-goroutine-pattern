package main
/**
按顺序输出任务其实就是做了一个有顺序的channel
 */
import (
  "fmt"
  "time"
)

func runTask(taskId, sleepTime int, ch chan string) {
  time.Sleep(time.Duration(sleepTime) * time.Second)
  ch <-fmt.Sprintf("任务id %d, 执行时间为 %d 秒", taskId, sleepTime)
  return
}

func main() {
  input := []int{3, 2, 1}
  chs := make([]chan string, len(input))
  startTime := time.Now()
  fmt.Println("Multirun start ...")

  for i, sleepTime := range input {
    chs[i] = make(chan string)
    go runTask(i, sleepTime, chs[i])
  }

  for _, ch := range chs {
    fmt.Println(<-ch)
  }

  endTime := time.Now()
  fmt.Printf("按顺序输出执行任务完成， 任务数为 %d， 耗时为 %s", len(input), endTime.Sub(startTime))
}
