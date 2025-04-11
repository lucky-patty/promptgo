package test

import (
  "testing"
  "github.com/lucky-patty/promptgo/selector"
)

func TestUnixRun_SelectSecondOption(t *testing.T) {
  keys := []string{"down", "enter"}
  i := 0

  mockReader := func() (string, error) {
    if i >= len(keys) {
      return "enter", nil
    }
    key := keys[i]
    i++
    return key, nil
  }

  options := []string{"first", "second", "third"}
  //selected, err := promptgo.Select("Choose:", options)
  selected, err := selector.UnixRunReader("Choose:", options, mockReader)
  if err != nil {
    t.Fatal(err)
  }

  if selected != "second" {
    t.Errorf("Expected 'second', got '%s'", selected)
  }
}
