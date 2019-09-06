package main

import (
  "fmt"
  "sync"
)

func doTask(wg *sync.WaitGroup, i int) {
  defer wg.Done()
  j := i * 2
  fmt.Println("task ", i , "done: ", j)
}

func main() {
  var wg sync.WaitGroup
  tasksNum := 10
  for i := 0; i < tasksNum; i++ {
    wg.Add(1)
    go doTask(&wg, i)
  }
  wg.Wait()
}
