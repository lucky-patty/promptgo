package selector 

import (
  "fmt"
  "runtime"
  "golang.org/x/term"
  "os"
  "os/signal"
  "syscall"

  "github.com/lucky-patty/promptgo/style"
)

func warnIfWindows() {
  if runtime.GOOS == "windows" {
    fmt.Println("⚠️  Your terminal may not fully support arrow key navigation.")
		fmt.Println("👉  Recommended: use Windows Terminal, WSL, or Git Bash for best experience.")
  }
}

func UnixRun(prompt string, options []string) (string, error) {
  
  // Try raw mode here
  fd := int(os.Stdin.Fd())
  oldState, err := term.MakeRaw(fd)
  if err != nil {
    panic(err)
  }
  // Restore terminal
  defer term.Restore(fd, oldState)

  theme := style.Current

  // Restore on SIGINT
  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
  go func() {
    <-sigs
    term.Restore(fd, oldState)
    os.Stdout.Write([]byte("\x1b[0m\r\n")) // Reset style, newline cleanly
    fmt.Println("\n🛑 Exit with Ctrl+C — terminal restored.")
    os.Exit(0)
  }()

  
  cursor := 0
  read := readInputUnix()
  
  render := func() {
    os.Stdout.Write([]byte("\x1b[2J")) // Clear screen 
    os.Stdout.Write([]byte("\x1b[H")) // Move cursor to 0,0
    // \r\n is the key since \n can cause error
    os.Stdout.Write([]byte(theme.Prompt + prompt + theme.Reset + "\r\n"))
    for i, opt := range options {
      if i == cursor {
        os.Stdout.Write([]byte(theme.Selected + "> " + opt + theme.Reset + "\r\n"))
      } else {
        os.Stdout.Write([]byte(theme.Unselected + opt + theme.Reset + "\r\n"))
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
    case "exit":
      term.Restore(fd, oldState)
      fmt.Fprintln(os.Stderr, "👋 User exited with Ctrl+C")
      os.Exit(0)
    }
    render()
  }
}

func Run(prompt string, options []string) (string, error) {
  if runtime.GOOS == "windows" {
    return "", fmt.Errorf("Windows is not yet supported in this version")
  }
  return UnixRun(prompt, options) 
}
