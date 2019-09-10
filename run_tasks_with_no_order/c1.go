package main

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
  ch := make(chan string)
  startTime := time.Now()
  fmt.Println("Multirun start ...")

  for i, sleepTime := range input {
    go runTask(i, sleepTime, ch)
  }

  for range input {
    fmt.Println(<-ch)
  }

  endTime := time.Now()
  fmt.Printf("并行执行任务完成， 任务数为 %d， 耗时为 %s", len(input), endTime.Sub(startTime))
}
