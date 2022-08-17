package datastruct

import (
	"fmt"
)

type AnyDataExporter struct {
	Name      string
	Data AnyType
}

func NewAnyDataExporter(n string, d AnyType) (a AnyDataExporter) {
	a.Name = n
	a.Data = d
	return
}

func (a AnyDataExporter) GetVariableName() string {
	return a.Name
}

func (a AnyDataExporter) ExportToCode() string {
	return fmt.Sprintf("%v := %v", a.Name, ToRightValue(a.Data))
}
