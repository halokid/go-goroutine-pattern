package main

import (
  "fmt"
  "sync"
  "time"
)

var wg sync.WaitGroup
var lk sync.RWMutex

func main() {
  m := make(map[int]int)
  for i := 0; i < 100; i++ {
    wg.Add(1)

    go func() {
      defer wg.Done()
      time.Sleep(10 * time.Second)
      lk.Lock()
      m[i] = i
      lk.Unlock()
    }()
  }

  wg.Wait()
  fmt.Println("run here........")
}
