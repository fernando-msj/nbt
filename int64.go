package nbt

import (
	"bytes"
	"unsafe"
)

func writeInt(value int64, buff *bytes.Buffer) {
	buff.Write(writeTypeHeader(TInt, int64bytes(value)))
}

func int64bytes(val int64) []byte {
	return (*(*[8]byte)(unsafe.Pointer(&val)))[:]
}

func int64Frombytes(val []byte) int64 {
	return *(*int64)(unsafe.Pointer(&val[0]))
}

//Serialize makes the value ready for disk
func (n Int) Serialize() *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})
	writeInt(int64(n), buffer)
	return buffer
}

//Int is an int
type Int int64

//Bytes serializes the int to byte
func (n Int) Bytes() []byte {
	return int64bytes(int64(n))
}

//Type returns the NBT type
func (n Int) Type() TType {
	return TInt
}
