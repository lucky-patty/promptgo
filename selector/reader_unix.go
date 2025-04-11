// go:build !windows
// +build !windows

package selector

import (
  "os"
  "fmt"
)

func runWithInput(prompt string, options []string, read func() (string, error)) (string, error) {
  cursor := 0
  for {
    key, _ := read()
    switch key { 
    case "up":
      if cursor > 0 {
        cursor--
      }
    case "down":
      if cursor < len(options) -1 {
        cursor++
      }
    case "enter":
      return options[cursor], nil
    case "exit":
      return "", fmt.Errorf("cancelled")
    }
  }
}

func readInputUnix() func() (string, error) {
  return func() (string, error) {
    var buf [3]byte
    os.Stdin.Read(buf[:])
    switch {
    // Case for Ctr+C 
    case buf[0] == 3:
      return "exit", nil
    case buf[0] == 13:
      return "enter", nil
    case buf[0] == 27 && buf[1] == 91:
      switch buf[2] {
      case 'A':
        return "up", nil 
      case 'B':
        return "down", nil
      }
    } 
    return "", nil
  }
}

