package util

import "testing"

func TestRandInt(t *testing.T) {
	min := 10
	max := 10000
	for i := 0; i < min; i++ {
		for j := max; j >= min; j-- {
			r := RandInt(i, j)
			//t.Log("RandInt min", i, "max:", j, "=", r)
			if r < i || r > j {
				t.Error("RandInt err")
			}
		}
	}
}

func TestRandIntSlice(t *testing.T) {
	for i := 0; i < 100; i++ {
		arr := RandIntSlice(10000, 6)
		if len(arr) != 6 {
			t.Error("RandIntSlice len err")
		}

		if len(sliceUnique(arr)) != len(arr) {
			t.Error("RandIntSlice has rep val")
		}

		//t.Log(arr)
		for _, v := range arr {
			if v > 10000 {
				t.Error("RandIntSlice val err")
			}
		}
	}
}

func BenchmarkRandInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandInt(10, 10000)
	}
}

func sliceUnique(s []int) []int {
	mp := make(map[int]int)
	for _, v := range s {
		mp[v] = 0
	}

	rs := []int{}
	for k, _ := range mp {
		rs = append(rs, k)
	}
	return rs
}
