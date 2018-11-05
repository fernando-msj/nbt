# NBT

Nammed binary tags (dead simple and error-free serialization format).
Note: this has absolutelly nothing to do with minecraft NBT format.

#### Just...Why???

* Message pack and others cannot map to non-concrete types.
* More cusomisable to suit my needs.
* Dynamic (deeply nested maps and arrays with with mixed datatypes)
* If you dont need an schema-free, flexible, _fast_ and dynamic, yet simple and small serialization method, nbt probably isn't what you are looking for.

#### Example

```go
package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"github.com/fernando-msj/nbt"
)

func main() {
	
	//An "complex" map
	data := nbt.Map{
		"string": nbt.String("Example"),
		"int": nbt.Int(10),
		"float": nbt.Float(3.14),
		"byte_array": nbt.ByteArray{0xa, 0xa, 0xa, 0xa, 0xa, 0xa},
		"array": nbt.Array{ nbt.Array{nbt.Int(10), nbt.Int(10)}, nbt.Array{nbt.Int(30) },
		"map": nbt.Map{ "X": nbt.Int(10), "Y": nbt.Int(10), "Z": nbt.Int(10) },
	}

	//Convert the data structure to an []byte
	nbtBytes := data.Serialize()

	//Write the binary to disc
	ioutil.WriteFile("test", nbtBytes.Bytes(), 0644)

	//Read again
	fileBytes, _ := ioutil.ReadFile("test")
	buffer := bytes.NewBuffer(fileBytes)

	//Parse the binary data
	parsed, err := nbt.ParseNBT(buffer)
	if err != nil { panic(err) }

	//Convert back to an nbt.Map
	newData, ok := parsed.(nbt.Map)
        if !ok { panic(err) }

	//Access the data
	fmt.Println(newData["string"])
	fmt.Println(newData["int"])
	fmt.Println(newData["float"])
	fmt.Println(newData["byte_array"])
	fmt.Println(newData["array"])
        fmt.Println(newData["map"])
}

```

#### Benchmark
```
goos: linux
goarch: amd64
BenchmarkArrayEmpty-8                   50000000                29.1 ns/op
BenchmarkArrayIntsBytes-8               50000000                29.2 ns/op
BenchmarkArrayIntsSerialize-8            3000000               427 ns/op
BenchmarkArrayComplexBytes-8            50000000                29.6 ns/op
BenchmarkArrayComplexSerialize-8         3000000               443 ns/op
BenchmarkSerializeMap-8                  1000000              1026 ns/op
BenchmarkBytesMap-8                      3000000               552 ns/op
BenchmarkSerializeComplexMap-8           1000000              1036 ns/op
BenchmarkBytesComplexMap-8               3000000               555 ns/op
```

#### Encoding
```
Type:      | Format:
------------------------------------------------------------------------
int        | type_header: 1 byte, value: 8 bytes
float      | type_header: 1 byte, value: 8 bytes
string     | type_header: 1 byte, byte_length: 8 bytes, string data: byte_length * 1 byte

byte_array | type_header: 1 byte, array_length: 8 bytes, value: array_length *  1 byte
array      | type_header: 1 byte, item_count: 8 bytes, data: item_count * (int|float|string|byte_array|array|map)
map        | type_header: 1 byte, key_count: 8 bytes,  data: item_count * (string, (int|float|string|byte_array|array|map))

Type headers:
int        = 0b00000001
float      = 0b00000010
string     = 0b00000011
byte_array = 0b00000100
array      = 0b00000101
map        = 0b00000110
```

#### TODO:
* Parallel serialization of "complex" maps (define complex)
* Support all number types
* Compact integer arrays (`nbt.Uint8Array{0, 1, 2}` instead of `nbt.IntArray{nbt.Int(0), nbt.Int(1), nbt.Int(2)}`)
* Compress type headers (Using Varints)
* Compress final buffer with gzip (package `"compress/gzip"`)
* Add more tests and benchmarks
* Make map keys an different type of string (type header not needed as map keys can only be strings)
* Map NBTs to structs using json tags
* Add API for custom encoders and decoders
