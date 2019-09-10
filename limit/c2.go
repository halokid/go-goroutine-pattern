package main

import (
  "fmt"
  "time"
)

func Run(taskId, sleepTime, timeout int, ch chan string) {
  chRun := make(chan string)
  go run(taskId, sleepTime, chRun)

  select {
  case re := <-chRun:
    ch <-re
  case <-time.After(time.Duration(timeout) * time.Second):
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
  chLimit := make(chan bool, 1)
  chs := make([]chan string, len(input))

  limitFunc := func(chLimit chan bool, ch chan string, taskId, sleepTime, timeOut int) {
    Run(taskId, sleepTime, timeOut, ch)
    <-chLimit
  }

  startTime := time.Now()
  fmt.Println("Multirun start ...")
  for i, sleepTime := range input {
    //chs[i] = make(chan string, 1)
    chs[i] = make(chan string)
    // 如果下面的gor func不执行的话，则会阻塞在这里
    chLimit <-true
    go limitFunc(chLimit, chs[i], i, sleepTime, timeOut)
  }

  for _, ch := range chs {
    fmt.Println(<-ch)
  }

  endTime := time.Now()
  fmt.Printf("按顺序输出执行任务完成， 任务数为 %d， 耗时为 %s", len(input), endTime.Sub(startTime))
}











