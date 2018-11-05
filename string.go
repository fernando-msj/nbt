package nbt

import (
	"bytes"
)

func stringBytes(str string) []byte {
	data := []byte(str)
	return append(int64bytes(int64(len(data))), data...)
}

func stringFromBytes(strBytes []byte) string {
	len := int64Frombytes(strBytes[:8])
	return string(strBytes[8 : 8+len])
}

func writeString(value string, buff *bytes.Buffer) {
	buff.Write(writeTypeHeader(TString, stringBytes(value)))
}

//String is an string
type String string

//Bytes serializes the string to byte
func (n String) Bytes() []byte {
	return stringBytes(string(n))
}

//Type returns the NBT type
func (n String) Type() TType {
	return TString
}

//Serialize makes the value ready for disk
func (n String) Serialize() *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})
	writeString(string(n), buffer)
	return buffer
}
