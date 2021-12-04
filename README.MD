<div align="center">
  <h1>
    <img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GO_PARIS.png" width="192" /><br />
    Abstract
  </h1>
</div>

If you want to be more clear about what your function is doing there are a couple of things you can do on your code. 

## Named Return

This is a really cool thing in Go:

When declaring your function, instead of just writing the return type, you can return a variable:

```go
func addTwoNumbers(a number, b: number) (sum number) {
  return a + B
}
```

In this case, `sum` adds a little bit more to our context.


## Go Docs

Another thing you can do is to use Go Docs which is simple as writing a comment, literally:

```go
// This comment will be your function doc.
// When hover the function name on your IDE or Editor
// This Doc will be highlighted.
func addTwoNumbers(a number, b: number) (sum number) {
  return a + B
}
```