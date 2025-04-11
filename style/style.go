package style 

import "runtime"

type Theme struct {
  Prompt string
  Selected string
  Unselected string
  Reset string
}

func IsWindows() bool {
  return runtime.GOOS == "windows"
}

var Current Theme

func init() {
  if IsWindows() {
    Current = WindowsTheme
  } else {
    Current = UnixTheme
  }
}
