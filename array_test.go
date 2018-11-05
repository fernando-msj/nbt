package nbt

import (
	"testing"
)

var complexArray Array
var intArray Array

func init() {
	array := make(Array, 1024)
	for i := 0; i < 1024; i++ {
		array[i] = Int(i * i)
	}
}

func init() {
	arrSize := 255

	//Bunch of ints
	array1 := make(Array, arrSize)
	for i := 0; i < arrSize; i++ {
		array1[i] = Int(i * i)
	}

	//Bunch of floats
	array2 := make(Array, arrSize)
	for i := 0; i < arrSize; i++ {
		array2[i] = Float((i * i) / 3)
	}

	//Bunch of strings
	array3 := make(Array, arrSize)
	for i := 0; i < arrSize; i++ {
		array3[i] = String("Testing testing testing testing testing")
	}

	//Bunch of arrays byte arrays
	array4 := make(Array, arrSize)
	for i := 0; i < arrSize; i++ {
		array4[i] = ByteArray([]byte{
			0xB0, 0xA0, 0xD0, 0x00, 0xF0, 0xA0, 0xC0, 0xE0,
			0xB0, 0xA0, 0xD0, 0x00, 0xF0, 0xA0, 0xC0, 0xE0,
			0xB0, 0xA0, 0xD0, 0x00, 0xF0, 0xA0, 0xC0, 0xE0,
			0xB0, 0xA0, 0xD0, 0x00, 0xF0, 0xA0, 0xC0, 0xE0,
			0xB0, 0xA0, 0xD0, 0x00, 0xF0, 0xA0, 0xC0, 0xE0,
			0xB0, 0xA0, 0xD0, 0x00, 0xF0, 0xA0, 0xC0, 0xE0,
			0xB0, 0xA0, 0xD0, 0x00, 0xF0, 0xA0, 0xC0, 0xE0,
			0xB0, 0xA0, 0xD0, 0x00, 0xF0, 0xA0, 0xC0, 0xE0,
		})
	}

	//Bunch of maps
	array5 := make(Array, arrSize)
	for i := 0; i < arrSize; i++ {
		array5[i] = Map{
			"a": Int(0),
			"b": Float(0),
			"c": String("Test"),
			"d": Array{Int(0), Float(0), String("Test")},
			"e": Map{
				"a": Int(0),
				"b": Float(0),
				"c": String("Test"),
				"d": Array{Int(0), Float(0), String("Test")},
			},
		}
	}

	complexArray := Array{}
	complexArray = append(complexArray, array1)
	complexArray = append(complexArray, array2)
	complexArray = append(complexArray, array3)
	complexArray = append(complexArray, array4)
	complexArray = append(complexArray, array5)
}

//BencArrayIntsBytes tests the conversion speed to bytes of an array with ints
func BenchmarkArrayEmpty(b *testing.B) {
	arr := Array{}
    for i := 0; i < b.N; i++ {
        arr.Bytes()
	}
}

//BencArrayIntsBytes tests the conversion speed to bytes of an array with ints
func BenchmarkArrayIntsBytes(b *testing.B) {
    for i := 0; i < b.N; i++ {
        intArray.Bytes()
	}
}

//BencArrayIntsSerialize tests the serialization speed an array with ints
func BenchmarkArrayIntsSerialize(b *testing.B) {
    for i := 0; i < b.N; i++ {
        intArray.Serialize()
    }
}

//BencArrayComplexBytes tests the conversion speed to bytes of an complex array
func BenchmarkArrayComplexBytes(b *testing.B) {
    for i := 0; i < b.N; i++ {
        complexArray.Bytes()
	}
}

//BencArrayComplexSerialize tests the serialization speed an complex array
func BenchmarkArrayComplexSerialize(b *testing.B) {
    for i := 0; i < b.N; i++ {
        complexArray.Serialize()
    }
}