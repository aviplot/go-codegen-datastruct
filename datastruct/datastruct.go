package datastruct

type DataStruct interface {
	GetVariableName() string // Output variable name
	ExportToCode() string    // Get the code itself
}

func NewDataStruct(t DataType, d AnyType, name string) (ds DataStruct, e error) {
	//switch t {
	//case ArrayDataType:
	//	k := reflect.TypeOf(d).Kind()
	//	if !(k == reflect.Array || k == reflect.Slice) {
	//		// Error, trying to create an array from not matched type.
	//		e = errors.New("not matched type (array/slice)")
	//		return
	//	}
	//
	//	//log.Printf("BaseType: %v", baseType)
	//	ds = NewArray(name, d)
	//	return
	//
	//case MapDataType:
	//	k := reflect.TypeOf(d).Kind()
	//	if !(k == reflect.Map) {
	//		// Error, trying to create a map from not matched type.
	//		e = errors.New("not matched type (map)")
	//		return
	//	}
	//	ds = NewMap(name, d)
	//	return
	//}
	ds = NewAnyDataExporter(name, d)
	return
}
