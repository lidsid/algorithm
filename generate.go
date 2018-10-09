package algorithm

import (
	"math/rand"
	"time"
)

/*生成随机数组*/
func Generate(start, end, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}
	nums := make([]int, 0, count)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		num := r.Intn(end-start) + start
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
