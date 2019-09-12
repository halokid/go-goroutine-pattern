package main

import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)

func doWork(job int) {
  time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
  fmt.Println("i am doing job", job)
}

func main() {
  workers := 3
  //jobs := make(chan int)
  jobs := make(chan int, 10)

  wg := &sync.WaitGroup{}
  wg.Add(workers)
  for i := 0; i < workers; i++ {
    go func() {
      defer wg.Done()
      for j := range jobs {
        doWork(j)
      }
    }()
  }


  for i := 0; i < 50; i++ {
    fmt.Println("get job", i)
    jobs <-i
  }
  close(jobs)

  wg.Wait()
  fmt.Println("all jobs done...")
}