package main

import (
  "fmt"
  "time"
)

func Run(taskId, sleepTime, timeOut int, ch chan string) {
  chRun := make(chan string)
  go run(taskId, sleepTime, chRun)

  select {
  case re := <-chRun:
    ch <- re
  case <-time.After(time.Duration(timeOut) * time.Second):
    re := fmt.Sprintf("任务id %d, 超时...", taskId)
    ch <-re
  }
}

func run(taskId, sleepTime int, ch chan string) {
  time.Sleep(time.Duration(sleepTime) * time.Second)
  ch <-fmt.Sprintf("任务id %d, 执行时间为 %d 秒", taskId, sleepTime)
  return
}

func main() {
  input := []int{3, 2, 1}
  timeOut := 2
  chs := make([]chan string, len(input))
  startTime := time.Now()
  fmt.Println("Multirun start ...")

  for i, sleepTime := range input {
    chs[i] = make(chan string)
    go Run(i, sleepTime, timeOut, chs[i])
  }

  for _, ch := range chs {
    fmt.Println(<-ch)
  }

  endTime := time.Now()
  fmt.Printf("按顺序输出执行任务，设置gor超时完成， 任务数为 %d， 耗时为 %s", len(input), endTime.Sub(startTime))
}






