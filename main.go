package main

import (
  "fmt"
  "github.com/kimpettersen/goCI/collector.go"
)

func main() {
  fmt.Println("Starting CI-server")

  c := collector.Collect('abc')
  fmt.Println("Got a result")
  fmt.Println(c)
}