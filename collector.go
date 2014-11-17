package main

import (
  "fmt"
  "os"
  "os/exec"
  "time"
  // "bufio"
  "path/filepath"
  "log"
  "strconv"
)

type Collector struct {
  path string
  url string
  id string
  name string
}

func Collect(name string, url string) Collector {

  uxTime := time.Now().Unix()
  t := strconv.FormatInt(uxTime, 10)

  name = name + "-" + t
  collector := Collector{ url: url, name: name }
  collector.url = url
  collector.path = generatePath(name)

  Clone(&collector)
  GetDependencies(&collector)
  RunTests()

  return collector
}

func generatePath(name string) string {
  currentPath, err := filepath.Abs(".")

  if err != nil {
    log.Println("Can't find current path ???")
  }

  _, err = os.Stat(currentPath + "/builds")

  if err != nil {
    err = os.Mkdir(currentPath + "/builds", 0777)
    if err != nil {
      fmt.Println("failed to create builds folder")
    }
  }

  ppath := currentPath + "/builds/" + name + "/"

  err = os.Mkdir(ppath, 0777)

  if err != nil {
    log.Println("Path already exists")
  }
  return ppath
}

func Clone(collector *Collector) {
  cmd := exec.Command("git", "clone", collector.url, collector.path)

  fmt.Println("Cloning:", collector.url)
  err := cmd.Run()

  if err != nil {
    log.Fatal("Error cloning repository: ", err)
  }

  fmt.Println("done")
}

func GetDependencies(collector *Collector) {
  cmd := exec.Command("npm", "--prefix", collector.path, "install", collector.path)
  fmt.Println(cmd)
  fmt.Println("Installing NPM dependencies")
  err := cmd.Run()

  if err != nil {
    log.Fatal("Error installing dependencies: ", err)
  }

  fmt.Println("done")
}

func RunTests() {
  cmd := exec.Command("npm", "test")
  fmt.Println("Running tests")
  err := cmd.Run()

  if err != nil {
    log.Fatal("Test failed: ", err)
  }

  fmt.Println("done")
}

// Move this to separate file

func main() {
  fmt.Println("Starting CI-server")

  Collect("math", "https://github.com/kimpettersen/goCI-test-project.git")


}