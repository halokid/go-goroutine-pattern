package main

import (
  "fmt"
  "time"
)

func Run(taskId, sleepTime int, ch chan string) {
  chRun := make(chan string)
  go run(taskId, sleepTime, chRun)
}

func run(taskId, sleepTime int, ch chan string) {
  time.Sleep(time.Duration(sleepTime) * time.Second)
  ch <-fmt.Sprintf("任务id %d, 执行时间为 %d 秒", taskId, sleepTime)
  return
}
