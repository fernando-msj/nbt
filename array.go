package nbt

import (
	"bytes"
)

func writeArray(data Array, buff *bytes.Buffer) {
	buff.Write(writeTypeHeader(data.Type(), data.Bytes()))
}

//Array is an array
type Array []TValue

//Bytes serializes the array to byte
func (n Array) Bytes() []byte {
	dataSection := []byte{}
	for _, value := range n {
		dataSection = append(dataSection, writeTypeHeader(value.Type(), value.Bytes())...)
	}
	headerSection := int64bytes(int64(len(n)))
	return append(headerSection, dataSection...)
}

//Type returns the NBT type
func (n Array) Type() TType {
	return TArray
}

//Serialize makes the value ready for disk
func (n Array) Serialize() *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})
	writeArray(n, buffer)
	return buffer
}
