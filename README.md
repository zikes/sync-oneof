# sync-oneof

`OneOf` will return true as soon as the first function argument returns true.
If no functions return true, `OneOf` will return false.

```go
funcs := []func() bool {
  func() bool {
    // Do some work
    return true
  },
  func() bool {
    // Do some work
    return false
  },
}

if oneof.OneOf(funcs...) {
  // At least one function returned true
  fmt.Println("true")
} else {
  // No functions return true
  fmt.Println("false")
}
```

Functions must have the signature `func() bool` but you can use closures to
inject work data into them:

```go
funcs := []func() bool{}
times := []time.Duration{
  100 * time.Millisecond,
  200 * time.Millisecond,
  50 * time.Millisecond,
  300 * time.Millisecond,
  1 * time.Second,
  500 * time.Millisecond,
}

for _, t := range times {
  t := t
  funcs = append(funcs, func() bool {
    select{
    case <- time.After(t):
      // Wait allotted time
      return false
    case <- time.After(600 * time.Millisecond):
      // Timeout
      return true
    }
  })
}

if OneOf(funcs...) {
  fmt.Println("timeout reached")
} else {
  fmt.Println("all functions completed successfully")
}

```
