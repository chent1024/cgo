package util

import (
	"math/rand"
	"time"
)

// 随机一个区间整数
func RandInt(min, max int) int {
	if min > max {
		return 0
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// 随机一个定长数组
func RandIntSlice(max, length int) []int {
	if length > max {
		return nil
	}

	var rs []int
	rand.Seed(time.Now().UnixNano())
	tmp := rand.Perm(max)
	rs = tmp[:length]
	return rs
}
