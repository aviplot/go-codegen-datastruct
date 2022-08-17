package datastruct

import (
	"fmt"
)

type AnyMap map[AnyType]AnyType

type Map struct {
	Name    string
	BaseMap AnyType
}

func NewMap(n string, bm AnyType) (m Map) {
	m.Name = n
	m.BaseMap = bm
	return
}

func (m Map) GetVariableName() string {
	return m.Name
}

func (m Map) ExportToCode() string {
	return fmt.Sprintf("%v := %v", m.Name, ToRightValue(m.BaseMap))
}
