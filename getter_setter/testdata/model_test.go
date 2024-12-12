package testdata

import (
	"fmt"
	"testing"
)

type Model2 struct {
	Model1
}

func (m *Model2) GetName() string {
	return "111"
}

func t2(model1 Model1Iface) {
	fmt.Println(model1.GetName())
}

func TestName(t *testing.T) {
	t2(&Model1{})
	t2(&Model2{})
}
