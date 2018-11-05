package nbt

import (
	"bytes"
	"encoding/binary"
	"math"
)

func float64bytes(float float64) []byte {
	val := make([]byte, 8)
	binary.LittleEndian.PutUint64(val, math.Float64bits(float))
	return val
}

func float64frombytes(bytes []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(bytes))
}

func writeFloat(value float64, buff *bytes.Buffer) {
	buff.Write(writeTypeHeader(TFloat, float64bytes(value)))
}

//Float is an float
type Float float64

//Bytes serializes the float to byte
func (n Float) Bytes() []byte {
	return float64bytes(float64(n))
}

//Type returns the NBT type
func (n Float) Type() TType {
	return TFloat
}

//Serialize makes the value ready for disk
func (n Float) Serialize() *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})
	writeFloat(float64(n), buffer)
	return buffer
}
