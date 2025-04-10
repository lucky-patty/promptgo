package promptgo

import (
  "github.com/lucky-patty/promptgo/selector"
)

func Select(prompt string, options []string) (string, error) {
  return selector.Run(prompt, options)
}
