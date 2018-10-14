package goStore

import (
	"errors"
	"encoding/binary"
)

type DataType int
const(
	INT8 DataType = iota
	INT16
	INT32
	INT64
	UINT8
	UINT16
	UINT32
	UINT64
	STRING
)
var(
	UnknownDataTypeError = errors.New("Unknown data type")
	MismatchTypeAndValueError = errors.New("Type and value mismatch")
)
var(
	toDataFunc=map[DataType]func(dataType DataType,value interface{})[]byte{
		INT8: func(dataType DataType, value interface{}) []byte {
			if v,ok:=value.(int8);ok{
				return []byte{byte(v)}
			}else{
				panic(MismatchTypeAndValueError)
			}
		},
		INT16: func(dataType DataType, value interface{}) []byte {
			if v,ok:=value.(int16);ok{
				var bs = make([]byte,2)
				binary.BigEndian.PutUint16(bs,uint16(v))
				return bs
			}else{
				panic(MismatchTypeAndValueError)
			}
		},
		INT32: func(dataType DataType, value interface{}) []byte {
			if v,ok:=value.(int32);ok{
				var bs = make([]byte,4)
				binary.BigEndian.PutUint32(bs,uint32(v))
				return bs
			}else{
				panic(MismatchTypeAndValueError)
			}
		},
		INT64: func(dataType DataType, value interface{}) []byte {
			if v,ok:=value.(int64);ok{
				var bs = make([]byte,8)
				binary.BigEndian.PutUint64(bs,uint64(v))
				return bs
			}else{
				panic(MismatchTypeAndValueError)
			}
		},
		UINT8: func(dataType DataType, value interface{}) []byte {
			if v,ok:=value.(uint8);ok{
				return []byte{v}
			}else{
				panic(MismatchTypeAndValueError)
			}
		},
		UINT16: func(dataType DataType, value interface{}) []byte {
			if v,ok:=value.(uint16);ok{
				var bs = make([]byte,2)
				binary.BigEndian.PutUint16(bs,v)
				return bs
			}else{
				panic(MismatchTypeAndValueError)
			}
		},
		UINT32: func(dataType DataType, value interface{}) []byte {
			if v,ok:=value.(uint32);ok{
				var bs = make([]byte,4)
				binary.BigEndian.PutUint32(bs,v)
				return bs
			}else{
				panic(MismatchTypeAndValueError)
			}
		},
		UINT64: func(dataType DataType, value interface{}) []byte {
			if v,ok:=value.(uint64);ok{
				var bs = make([]byte,8)
				binary.BigEndian.PutUint64(bs,v)
				return bs
			}else{
				panic(MismatchTypeAndValueError)
			}
		},
		STRING: func(dataType DataType, value interface{}) []byte {
			if v,ok:=value.(string);ok{
				return []byte(v)
			}else{
				panic(MismatchTypeAndValueError)
			}
		},
	}
)
func dataTypeToData(dataType DataType,value interface{})[]byte{
	if f:=toDataFunc[dataType];f!=nil{
		return f(dataType,value)
	}else{
		panic(UnknownDataTypeError)
	}
}