# getter_setter

To implement `extend` function


```go
package main

type Model1 struct {
  // some fields...
}

// getters & setters implements of fields

type Model1Iface interface {
	// getters & setters declears of fields
}

type Model2 struct {
	Model1
}

func main(){
  t1(&Model2{}) // no!
  t2(&Model2{}) // its no problem!
}

func t1(m Model1) { // use the struct
	fmt.Println(m.GetName())
}

func t2(m Model1Iface) { // use the interface
	fmt.Println(m.GetName())
}
```
