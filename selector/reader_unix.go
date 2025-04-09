// go:build !windows
// +build !windows

package selector

import (
  "golang.org/x/term"
  "os"
  "os/signal"
  "syscall"
)

func readInputUnix() func() (string, error) {
  fd := int(os.Stdin.Fd())
  oldState, _ := term.MakeRaw(fd)

  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, os.Interrupt)
  go func() {
    <-sigs
    term.Restore(fd, oldState)
    os.Exit(1)
  }()

  return func() (string, error) {
    var buf [3]byte
    os.Stdin.Read(buf[:])

    switch {
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
