package main

import (
  "fmt"
)

func doTask(i int, tasksResStream chan int, taskNum int, done chan int) {
  j := i * 2
  tasksResStream <-j
  fmt.Println("task ", i , "done: ", j)
  if i == 10 {
    //defer close(tasksResStream)
    //close(tasksResStream)
    done <-1
  }
}


func main() {
  taskRes := make(map[int]int)
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
      taskRes[j/2] = j          // 以流运算去代替共享内存
      fmt.Println("task stream read ------ ", j)
    case <-done:
      //close(tasksResStream)
      fmt.Println("taskRes ------------- ", taskRes)
      return
    }
  }

}
