package Tools

import (
	"math"
	"math/rand"
	"time"
)

/*
	生成定长随机小数
	left 左
	right 右
	bit 长度
*/
func RandFloat(left int, right int, bit int) float32 {
	// 容错
	if left > right {
		var tmp = right
		right = left
		left = tmp
	}
	var res float32 = 0
	//创建随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < bit; i++ {
		res = res*10 + float32(rand.Intn(10))
	}
	res /= float32(math.Pow10(bit))
	res += float32(rand.Intn(right-left) + left)
	return res
}
