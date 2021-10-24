package grains

/*
左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。
其功能把"<<"左边的运算数的各二进位全部左移若干位，
由"<<"右边的数指定移动的位数，高位丢弃，低位补0。
*/
import "errors"

func Square(input int) (uint64, error) {
	if input > 64 || input <= 0 {
		return 0, errors.New("invalid input")
	}
	return uint64(1 << (input - 1)), nil
}

func Total() (sum uint64) {
	return uint64(1<<64 - 1) // 等比数列计算公式求得。
}
