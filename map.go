package nbt

import (
	"bytes"
)

func writeMap(data Map, buff *bytes.Buffer) {
	buff.Write(writeTypeHeader(data.Type(), data.Bytes()))
}

//Map is an map
type Map map[String]TValue

//Bytes serializes the map to byte
func (n Map) Bytes() []byte {
	data := int64bytes(int64(len(n)))
	for key, value := range n {
		data = append(data, key.Bytes()...)
		data = append(data, writeTypeHeader(value.Type(), value.Bytes())...)
	}
	return data
}

//Type returns the NBT type
func (n Map) Type() TType {
	return TMap
}

//Serialize makes the value ready for disk
func (n Map) Serialize() *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})
	writeMap(n, buffer)
	return buffer
}
