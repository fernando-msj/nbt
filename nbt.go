package nbt

import (
	"fmt"
	"bytes"
	"errors"
)

//  Type:      | Format:
//--------------------------------------------------------------------------
//             |
//  int        | type: 00000001 value: 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000
//  float      | type: 00000010 value: 00000000 00000000 00000000 00000000
//  string     | type: 00000011 str_len: 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 value: [*str_len: 00000000]

//  byte_array | type: 00000100 arr_size: 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 value: [*arr_size: 00000000]
//  array      | type: 00000101 arr_size: 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 value: [*arr_size: (int, float, string, byte_array, array, map)]
//  map        | type: 00000110 key_count: 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 [*key_count: {key: str_len: 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 value: [*str_len: 00000000], value: (int, float, string, byte_array, array, map)}  ]

//TType is used for writting the header
type TType byte

//All supported types
const (
	TInt       TType = 1
	TFloat     TType = 2
	TString    TType = 3
	TByteArray TType = 4
	TArray     TType = 5
	TMap       TType = 6
)

//TValue represents an serializable value
type TValue interface {
	Bytes() []byte
	Type() TType
	Serialize() *bytes.Buffer
}

func writeTypeHeader(xtype TType, data []byte) []byte {
	return append([]byte{byte(xtype)}, data...)
}

func sizeBytes() []byte {
	return []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
	}
}

//ParseNBT returns the parsed buffer
func ParseNBT(buff *bytes.Buffer) (TValue, error) {
	typeHeader, err := buff.ReadByte()
	
	if err != nil {
		return nil, errors.New("Failed to read type header: " + err.Error())
	}

	var parseError error
	var nbtValue TValue

	switch TType(typeHeader) {

	case TInt:
		p := sizeBytes()
		_, err = buff.Read(p)
		if err != nil {
			parseError = errors.New("Failed to read int: " + err.Error())
			break;
		}
		nbtValue = Int(int64Frombytes(p))
		break
	case TFloat:
		p := sizeBytes()
		_, err = buff.Read(p)
		if err != nil {
			parseError = errors.New("Failed to read float: " + err.Error())
			break;
		}
		nbtValue = Float(float64frombytes(p))
		break
	case TString:
		p := sizeBytes()
		_, err = buff.Read(p)
		if err != nil {
			parseError = errors.New("Failed to read string header: " + err.Error())
			break;
		}
		len := int64Frombytes(p)
		data := make([]byte, len)
		_, err = buff.Read(data)
		if err != nil {
			parseError = errors.New("Failed to read string contents: " + err.Error())
			break;
		}
		nbtValue = String(string(data))
		break
	case TByteArray:
		p := sizeBytes()
		_, err = buff.Read(p)
		if err != nil {
			parseError = errors.New("Failed to read byte array header: " + err.Error())
			break;
		}
		len := int64Frombytes(p)
		data := make([]byte, len)
		_, err = buff.Read(data)
		if err != nil {
			parseError = errors.New("Failed to read byte array header: " + err.Error())
			break;
		}
		nbtValue = ByteArray(data)
		break
	case TArray:
		p := sizeBytes()
		_, err = buff.Read(p)
		if err != nil {
			parseError = errors.New("Failed to array header: " + err.Error())
			break;
		}
		len := int64Frombytes(p)
		nbtArray := make(Array, len)
		for i := int64(0); i < len; i++ {
			val, err := ParseNBT(buff)
			if err != nil {
				parseError = errors.New(fmt.Sprintf("Nested parse error: failed to read value of item %v in array:", i)  + err.Error())
				break;
			}
			nbtArray[i] = val
		}
		nbtValue = nbtArray
		break
	case TMap:
		p := sizeBytes()
		_, err := buff.Read(p)
		if err != nil {
			parseError = errors.New("Failed to nbt map header: " + err.Error())
			break;
		}
		len := int64Frombytes(p)
		nbtMap := Map{}
		for i := int64(0); i < len; i++ {
			p := sizeBytes()
			_, err = buff.Read(p)
			if err != nil {
				parseError = errors.New(fmt.Sprintf("Nested parse error: failed to read key of item %v in nbt map:", i)  + err.Error())
				break;
			}
			len := int64Frombytes(p)
			data := make([]byte, len)
			buff.Read(data)
			val, err := ParseNBT(buff)
			if err != nil {
				parseError = errors.New(fmt.Sprintf("Nested parse error: failed to read value of item %v in nbt map:", string(data))  + err.Error())
				break;
			}
			nbtMap[String(string(data))] = val
		}
		nbtValue = nbtMap
		break
	}
	return nbtValue, parseError
}