package nbt

import (
	"bytes"
)

func writeByteArray(data []byte, buff *bytes.Buffer) {
	buff.Write(writeTypeHeader(TByteArray, byteArrayBytes(data)))
}

func byteArrayBytes(data []byte) []byte {
	return append(int64bytes(int64(len(data))), data...)
}

//ByteArray is an byte array
type ByteArray []byte

//Bytes serializes the byte array to byte
func (n ByteArray) Bytes() []byte {
	return byteArrayBytes([]byte(n))
}

//Type returns the NBT type
func (n ByteArray) Type() TType {
	return TByteArray
}

//Serialize makes the value ready for disk
func (n ByteArray) Serialize() *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})
	writeByteArray([]byte(n), buffer)
	return buffer
}
