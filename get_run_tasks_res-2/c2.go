package main

import (
  "fmt"
)

func doTask(i int, tasksResStream chan int, taskNum int, done chan int) {
  if i == taskNum {
    defer done <-1
  }
  j := i * 2
  tasksResStream <-j
  fmt.Println("task ", i , "done: ", j)
}


func main() {
  tasksResStream := make(chan int)
  done := make(chan int)
  taskNum := 10
  for i := 1; i <= taskNum; i++ {
    /**
    // fixme: 内存域污染陷阱
    go func() {
      doTask(i, tasksResStream)
    }()
    */
    go doTask(i, tasksResStream, taskNum, done)
  }


  // reading stream
  for {
    select {
    case j := <-tasksResStream:
      fmt.Println("task stream read ------ ", j)
    }
  }
}
