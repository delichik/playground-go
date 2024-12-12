# getter_setter

To implement `extend` function

## example

There is a function, now is `the_funtion1`, which needs to call the `A()`, and the function is allow any struct which include `P`.

```go
type P struct {
	msg1 string
}

func (p *P) A() string {
	return p.msg1
}

type T struct {
	*P
	msg2 string
}

func (t *T) A() string {
	return t.msg2
}

func the_function1(p *P) {
	fmt.Println(p.A())
}
```

We can only use like this:

```go
func main() {
	t := &T{
		P: &P{
			msg1: "1",
		},
		msg2: "2",
	}
	the_function1(t.P)
}
```
It prints `1`.

But we want to call `(t *T) A() string` instead of `(p *P) A() string`.

Use the generator, and write another `the_function`:

```go
func the_function2(p PIface) {
	fmt.Println(p.A())
}

func main() {
	t := &T{
		P: &P{
			msg1: "1",
		},
		msg2: "2",
	}
	the_function2(t)
}
```
It prints `2`!
