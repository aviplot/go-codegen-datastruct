package datastruct

import (
	"fmt"
)

type AnyArray []AnyType

type Array struct {
	Name      string
	BaseArray AnyType
}

func NewArray(n string, ba AnyType) (a Array) {
	a.Name = n
	a.BaseArray = ba
	return
}

func (a Array) GetVariableName() string {
	return a.Name
}

func (a Array) ExportToCode() string {
	return fmt.Sprintf("%v := %v", a.Name, ToRightValue(a.BaseArray))
}
