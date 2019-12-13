package util

import "testing"

func TestStr2Time(t *testing.T) {
	v, err := Str2Time("2019-12-13 00:00:00")
	//t.Log("2019-12-13 00:00:00 = ", v)
	if v != int64(1576166400) {
		t.Error("str2time err, ", err)
	}
}
