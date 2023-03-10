package main

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	t.Skip()
	arr := Array{
		Length: 0,
		Data:   map[uint]interface{}{},
	}

	arr.Append(1)
	res := arr.Get(0)
	len := arr.Length

	if res != 1 {
		t.Error("Error in append")
	}

	if len != 1 {
		t.Error("Error in length")
	}
}

func TestMap(t *testing.T) {
	t.Skip()
	arr := Array{
		Length: 0,
		Data:   map[uint]interface{}{},
	}

	arr.Append(1)
	arr.Append(2)

	res := arr.Map(mapTest)

	if res.Get(0) != 2 {
		t.Error("Error in append")
	}
}

func TestUnshift(t *testing.T) {
	t.Skip()
	arr := Array{
		Length: 0,
		Data:   map[uint]interface{}{},
	}

	arr.Append(1)
	arr.Append(2)
	arr.Unshift(3)

	if arr.Get(0) != 3 {
		t.Error("Error in unshift")
	}

	if arr.Length != 3 {
		t.Error("Error in unshift length")
	}
}

func TestPop(t *testing.T) {
	arr := Array{
		Length: 0,
		Data:   map[uint]interface{}{},
	}

	arr.Append(1)
	arr.Append(2)
	res := arr.Pop()

	if res != 2 {
		t.Error("Error in pop")
	}

	if arr.Length != 1 {
		t.Error("Error in pop length")
	}
}

func TestConcat(t *testing.T) {
	t.Skip()
	arr1 := Array{
		Length: 0,
		Data:   map[uint]interface{}{},
	}

	arr2 := Array{
		Length: 0,
		Data:   map[uint]interface{}{},
	}

	arr1.Append(1)
	arr1.Append(2)

	arr2.Append(3)
	arr2.Append(4)

	arr1.Concat(arr2)

	if arr1.Get(0) != 1 {
		t.Error("Error in concat 1")
	}

	if arr1.Get(1) != 2 {
		t.Error("Error in concat 2")
	}

	if arr1.Get(2) != 3 {
		t.Error("Error izn concat 3")
	}

	if arr1.Get(3) != 4 {
		t.Error("Error in concat 4")
	}

	if arr1.Length != 4 {
		t.Error("Error in concat length")
	}
}

func TestSplice(t *testing.T) {
	arr := Array{
		Length: 0,
		Data:   map[uint]interface{}{},
	}

	arr.Append(1)
	arr.Append(2)

	arr.Splice(2, 0, 3, 4, 5, 6)

	fmt.Println(arr)

	if arr.Get(0) != 1 {
		t.Error("Error in splice")
	}

	if arr.Length != 6 {
		t.Error("Error in splice length")
	}
}

func mapTest(value interface{}, index uint) interface{} {
	return value.(int) + 1
}
