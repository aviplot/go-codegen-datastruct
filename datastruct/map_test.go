package datastruct

import (
	"log"
	"testing"
)

func TestMap(t *testing.T) {

	d := map[int]string{
		1: "Hello",
		2: "World",
		3: "And",
		4: "More",
	}

	log.Printf("TestMap")
	ds, e := NewDataStruct(MapDataType, d, "DataVariableName")
	if e != nil {
		t.Fatalf("Error: %v", e)
	}
	code := ds.ExportToCode()

	//t.Logf(code)
	log.Printf(code)
}
