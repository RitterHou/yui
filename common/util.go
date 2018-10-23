package common

import (
	"bytes"
	"encoding/binary"
	"math"
)

// 字节的大小端
var end = binary.BigEndian

func Int2ByteArray(i int32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, end, i)
	return buf.Bytes()
}

func ByteArray2Int(buf []byte) int32 {
	var i int32
	binary.Read(bytes.NewReader(buf), end, &i)
	return i
}

func Float2ByteArray(f float32) []byte {
	buf := make([]byte, 4)
	bits := math.Float32bits(f)
	end.PutUint32(buf, bits)
	return buf
}

func ByteArray2Float(buf []byte) float32 {
	bits := end.Uint32(buf)
	return math.Float32frombits(bits)
}

func IsCompiled(data []byte) bool {
	if len(data) < 4 {
		return false
	}
	for i := 0; i < 4; i++ {
		if data[i] != Magic[i] {
			return false
		}
	}
	return true
}
