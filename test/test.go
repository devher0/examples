package main

import (
	"fmt"
	"unsafe"
)

var (
	ToBe      bool   = false
	MaxInt8   int8   = 1<<7 - 1
	Maxuint8  uint8  = 1<<8 - 1
	MaxInt16  int16  = 1<<15 - 1
	Maxuint16 uint16 = 1<<16 - 1
	MaxInt32  int32  = 1<<31 - 1
	Maxuint32 uint32 = 1<<32 - 1
	MaxInt64  int64  = 1<<63 - 1
	Maxuint64 uint64 = 1<<64 - 1
	MaxByte   byte   = 1<<8 - 1
	MaxRune   rune   = 1<<31 - 1
	MaxInt    int    = 1<<63 - 1
)

func main() {

	fmt.Printf("Тип: %T размер: %d байт значение: %v\n", ToBe, unsafe.Sizeof(ToBe), ToBe)
	fmt.Printf("Тип: %T размер: %d байт значение: %v\n", MaxInt8, unsafe.Sizeof(MaxInt8), MaxInt8)
	fmt.Printf("Тип: %T размер: %d байт значение: %v\n", Maxuint8, unsafe.Sizeof(Maxuint8), Maxuint8)
	fmt.Printf("Тип: %T размер: %d байт значение: %v\n", MaxInt16, unsafe.Sizeof(MaxInt16), MaxInt16)
	fmt.Printf("Тип: %T размер: %d байт значение: %v\n", Maxuint16, unsafe.Sizeof(Maxuint16), Maxuint16)
	fmt.Printf("Тип: %T размер: %d байт значение: %v\n", MaxInt32, unsafe.Sizeof(MaxInt32), MaxInt32)
	fmt.Printf("Тип: %T размер: %d байт значение: %v\n", Maxuint32, unsafe.Sizeof(Maxuint32), Maxuint32)
	fmt.Printf("Тип: %T размер: %d байт значение: %v\n", MaxInt64, unsafe.Sizeof(MaxInt64), MaxInt64)
	fmt.Printf("Тип: %T размер: %d байт значение: %v\n", Maxuint64, unsafe.Sizeof(Maxuint64), Maxuint64)
	fmt.Println("=================")
	fmt.Printf("Тип: %T размер: %d байт значение: %v\n", MaxByte, unsafe.Sizeof(MaxByte), MaxByte)
	fmt.Printf("Тип: %T размер: %d байт значение: %v\n", MaxRune, unsafe.Sizeof(MaxRune), MaxRune)
	fmt.Printf("Тип: %T размер: %d байт значение: %v\n", MaxInt, unsafe.Sizeof(MaxInt), MaxInt)

}
