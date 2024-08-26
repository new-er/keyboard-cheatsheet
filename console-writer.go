package main

import (
  "fmt"
)

var(
  writtenLines int
)
func ConsoleWriteLine(text string) {
  fmt.Println(text)
  writtenLines++
}

func ConsoleClearWrittenLines(){
  for i := 0; i < writtenLines; i++ {
    fmt.Print("\x1b[1A\x1b[2K")
  }
  writtenLines = 0
}
