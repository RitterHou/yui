package common

import (
	"bytes"
	"encoding/binary"
	"math"
)

func Int2ByteArray(i uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, i)
	return buf
}

func Float2ByteArray(f float32) []byte {
	i := math.Float32bits(f)
	return Int2ByteArray(i)
}

func ByteArray2Int(buf []byte) uint32 {
	var i uint32
	binary.Read(bytes.NewReader(buf), binary.BigEndian, &i)
	return i
}

func ByteArray2Float(buf []byte) float32 {
	i := binary.BigEndian.Uint32(buf)
	return math.Float32frombits(i)
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
