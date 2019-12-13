package util

import "testing"

func TestSubStr(t *testing.T) {
	str := "我是中国人■※βㄛp玊"
	sub := SubStr(str, 5, "")
	t.Log(sub)
	if sub != "我是中国人" {
		t.Error("SubStr limit err")
	}

	sub = SubStr(str, 5, "...")
	t.Log(sub)
	if sub != "我是中国人..." {
		t.Error("SubStr suffix err")
	}

	str = "我是※βㄛ中国人■※βㄛp玊"
	sub = SubStr(str, 5, "")
	t.Log(sub)
	if sub != "我是※βㄛ" {
		t.Error("SubStr limit err")
	}

	sub = SubStr(str, 5, "...")
	t.Log(sub)
	if sub != "我是※βㄛ..." {
		t.Error("SubStr suffix err")
	}
}
