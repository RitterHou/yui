package common

import (
	"bytes"
	"encoding/binary"
	"math"
)

func Int2ByteArray(i int32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i)
	return buf.Bytes()
}

func ByteArray2Int(buf []byte) int32 {
	var i int32
	binary.Read(bytes.NewReader(buf), binary.LittleEndian, &i)
	return i
}

func Float2ByteArray(f float32) []byte {
	buf := make([]byte, 4)
	bits := math.Float32bits(f)
	binary.LittleEndian.PutUint32(buf, bits)
	return buf
}

func ByteArray2Float(buf []byte) float32 {
	bits := binary.LittleEndian.Uint32(buf)
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
