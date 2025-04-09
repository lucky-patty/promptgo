package selector 

import (
  "fmt"
  "runtime"
)

func warnIfWindows() {
  if runtime.GOOS == "windows" {
    fmt.Println("⚠️  Your terminal may not fully support arrow key navigation.")
		fmt.Println("👉  Recommended: use Windows Terminal, WSL, or Git Bash for best experience.")
  }
}

func Run(prompt string, options []string) (string, error) {
  if runtime.GOOS == "windows" {
    return "", fmt.Errorf("Windows is not yet supported in this version")
  }
  cursor := 0
  read := readInputUnix()
  
  render := func() {
    fmt.Print("\x1b[2J\x1b[H") // clear screen 
    fmt.Println(prompt)
    for i, opt := range options {
      if i == cursor {
        fmt.Printf("\x1b[7m> %s\x1b[0m\n", opt)
      } else {
        fmt.Println(" " + opt)
      }
    }
  }

  render()

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
    }
    render()
  }
}
