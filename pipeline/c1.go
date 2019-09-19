package main

import (
  "context"
  "github.com/pkg/errors"
)

func lineListSource(ctx context.Context, lines ...string) (<-chan string, <-chan error, error) {
  if len(lines) == 0 {
    return nil, nil, errors.Errorf("no lines provided")
  }

  out := make(chan string)
  errc := make(chan error, 1)

  go func() {
    defer close(out)
    defer close(errc)
    for lineIdx, line := range lines {
      if line == "" {
        errc <- errors.Errorf("line %v is empty", lineIdx+1)
        return
      }

      select {
      case out <-line:
      case <-ctx.Done():
        return
      }
    }
  }()
  return out, errc, nil
}





