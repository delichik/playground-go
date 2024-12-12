package generator

import (
	"fmt"
	"reflect"
)

var typeOfStringer = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

type Target struct {
	Model any

	TableName string
}
