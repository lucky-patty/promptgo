// go:build !windows
// +build !windows

package selector

import (
  "os"
)

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

