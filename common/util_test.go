package common

import "testing"

func TestIntConvert(t *testing.T) {
	var i int32 = 23333
	j := ByteArray2Int(Int2ByteArray(i))
	if i == j {
		t.Log("pass")
	} else {
		t.Errorf("j is %d", j)
	}
}

func TestFloatConvert(t *testing.T) {
	var i float32 = 2.3333
	j := ByteArray2Float(Float2ByteArray(i))
	if i == j {
		t.Log("pass")
	} else {
		t.Errorf("j is %.10f", j)
	}
}
