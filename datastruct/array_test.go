package datastruct

import (
	"log"
	"testing"
)

func TestArray1(t *testing.T) {
	d := []string{
		"Hello",
		"World",
		"And with \"quot\"",
		"More",
	}

	ds, e := NewDataStruct(ArrayDataType, d, "DataVariableName")
	if e != nil {
		t.Fatalf("Error: %v", e)
	}
	code := ds.ExportToCode()
	//t.Logf(code)
	log.Printf(code)
}
func TestArray2(t *testing.T) {
	d := []int{
		1,
		2,
		-3,
		4,
	}

	ds, e := NewDataStruct(ArrayDataType, d, "DataVariableName")
	if e != nil {
		t.Fatalf("Error: %v", e)
	}
	code := ds.ExportToCode()
	//t.Logf(code)
	log.Printf(code)
}

func TestArray3(t *testing.T) {
	d := []float64{
		1.12341234121234245,
		2.45623452352,
		-3.34562342563255434,
		4.7634534234564323,
	}

	ds, e := NewDataStruct(ArrayDataType, d, "DataVariableName")
	if e != nil {
		t.Fatalf("Error: %v", e)
	}
	code := ds.ExportToCode()
	//t.Logf(code)
	log.Printf(code)
}

func TestArray4(t *testing.T) {
	d := []map[int]string{
		{1: "a", 2: "b"},
		{11: "aa", 222: "bbb"},
	}

	ds, e := NewDataStruct(ArrayDataType, d, "DataVariableName")
	if e != nil {
		t.Fatalf("Error: %v", e)
	}
	code := ds.ExportToCode()
	//t.Logf(code)
	log.Printf(code)
}

func TestArray5(t *testing.T) {

	type heads struct {
		Eyes  string
		Mouth int
		Nose  string
	}
	type person struct {
		MyName string
		age    int
		Head   heads
	}

	d := person{
		MyName: "myname",
		age:    16,
		Head: heads{
			Eyes:  "brown",
			Mouth: 4,
			Nose:  "Nooooosssssse",
		},
	}

	//ds, e := NewDataStruct(ArrayDataType, d, "DataVariableName")
	//if e != nil {
	//	t.Fatalf("Error: %v", e)
	//}
	//code := ds.ExportToCode()
	////t.Logf(code)
	log.Printf("Struct: %v", ToRightValue(d))
}
