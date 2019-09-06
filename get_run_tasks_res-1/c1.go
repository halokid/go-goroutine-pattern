package main

import (
  "fmt"
  "sync"
)

func doTask(wg *sync.WaitGroup, lk *sync.RWMutex, i int, taskRes *map[int]int) {
  defer wg.Done()
  j := i * 2
  lk.Lock()
  (*taskRes)[i] = j
  lk.Unlock()
  fmt.Println("task ", i , "done: ", j)
}


func main() {
  taskRes := make(map[int]int)
  var wg sync.WaitGroup
  var lk sync.RWMutex
  taskNum := 10
  for i := 0; i < taskNum; i++ {
    wg.Add(1)
    go doTask(&wg, &lk, i, &taskRes)
  }
  wg.Wait()
}
