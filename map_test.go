package nbt

import (
	"testing"
)

var complexMap Map

func init() {
	complexMap = Map{
		"name": String("Testing testing testing testing"),
		"id": Int(10000000),
		"position": Map{
			"x": Float(0),
			"y": Float(0),
			"z": Float(0),
		},
		"speed": Map{
			"x": Float(0),
			"y": Float(0),
			"z": Float(0),
		},
		"prexPos": Map{
			"x": Float(0),
			"y": Float(0),
			"z": Float(0),
		},
		"flags": Map{
			"canDoA": Int(0),
			"canDoB": Int(1),
			"canDoC": Int(0),
			"canDoD": Int(1),
			"canDoE": Int(1),
			"canDoF": Int(0),
		},
		"textureBitmap": ByteArray{
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
			0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0, 0xA0,
		},
		"attributes": Array{
			Map{
				"type": String("X"),
				"value": String("Y"),
				"multiplier": Int(10),
				"base": Float(0),
				"nested": Map{
					"A1": Map{ "x": Int(0), "y": Float(0) },
					"A2": Map{ "x": Int(0), "y": Float(0) },
					"A3": Map{ "x": Int(0), "y": Float(0) },
					"A4": Map{ "x": Int(0), "y": Float(0) },
				},
			},
			Map{
				"type": String("X"),
				"value": String("Y"),
				"multiplier": Int(10),
				"base": Float(0),
				"nested": Map{
					"A1": Float(0),
					"A2": Float(0),
					"A3": Float(0),
					"A4": Float(0),
				},
			},
			Map{
				"type": String("X"),
				"value": String("Y"),
				"multiplier": Int(10),
				"base": Float(0),
				"nested": Map{
					"A1": Int(0),
					"A2": Int(0),
					"A3": Int(0),
					"A4": Int(0),
				},
			},
			Map{
				"type": String("X"),
				"value": String("Y"),
				"multiplier": Int(10),
				"base": Float(0),
				"nested": Map{
					"A1": String("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
					"A2": String("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
					"A3": String("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
					"A4": String("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
				},
			},
			Map{
				"type": String("X"),
				"value": String("Y"),
				"multiplier": Int(10),
				"base": Float(0),
				"nested": Map{
					"A1": ByteArray{0xA,0xA,0xA,0xA,0xA,0xA,0xA,0xA,0xA,},
					"A2": ByteArray{0xA,0xA,0xA,0xA,0xA,0xA,0xA,0xA,0xA,},
					"A3": ByteArray{0xA,0xA,0xA,0xA,0xA,0xA,0xA,0xA,0xA,},
					"A4": ByteArray{0xA,0xA,0xA,0xA,0xA,0xA,0xA,0xA,0xA,},
				},
			},
		},
	}
}

//BenchmarkSerializeMap tests the serialization speed of an map with one string value
func BenchmarkSerializeMap(b *testing.B) {
	testMap := Map{ "value": String("test test test") }
    for i := 0; i < b.N; i++ {
        testMap.Serialize()
    }
}

//BenchmarkBytesMap tests the conversion to bytes speed of an map with one string value
func BenchmarkBytesMap(b *testing.B) {
	testMap := Map{ "value": String("test test test") }
    for i := 0; i < b.N; i++ {
        testMap.Bytes()
    }
}

//BenchmarkSerializeComplexMap tests the serialization speed of an map with one string value
func BenchmarkSerializeComplexMap(b *testing.B) {
	testMap := Map{ "value": String("test test test") }
    for i := 0; i < b.N; i++ {
        testMap.Serialize()
    }
}

//BenchmarkBytesComplexMap tests the conversion to bytes speed of an map with one string value
func BenchmarkBytesComplexMap(b *testing.B) {
	testMap := Map{ "value": String("test test test") }
    for i := 0; i < b.N; i++ {
        testMap.Bytes()
    }
}