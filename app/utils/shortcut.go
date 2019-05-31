package utils

import (
	"math"
	"strings"
)

var charset string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//GenerateShort 把10进制转换成62进制
func GenerateShort(num int64) string {
	var short []byte
	//进制转换规则
	//每次将整数部分除以62，余数为该位权上的数，而商继续除以62，余数又为上一个位权上的数，这个步骤一直持续下去，直到商为0为止，最后读数时候，从最后一个余数读起，一直到最前面的一个余数
	for {
		number := num % 62
		result := charset[number]

		short = append(short, result)
		num = num / 62
		if num == 0 {
			break
		}
	}
	//由于在上面for循环中没有按顺序处理转换后的字符串，所以此处需要倒转数组
	if len(short) > 1 {
		for begin, end := 0, len(short)-1; begin < end; begin, end = begin+1, end-1 {
			short[begin], short[end] = short[end], short[begin]
		}
	}

	return string(short)
}

//ParseShort 把62进制转化成10进制
func ParseShort(short string) int64 {

	var (
		num int64
		pos int
	)

	for index, len := 0, len(short); index < len; index++ {
		pos = strings.IndexAny(charset, short[index:index+1])
		num += int64(math.Pow(62, float64(len-index-1)) * float64(pos))
	}

	return num
}
