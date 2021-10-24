package secret

var decides = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

func Handshake(num uint) (output []string) {
	for _, str := range decides {
		if num%2 == 1 {
			output = append(output, str)
		}
		num /= 2
	}
	if num%2 == 1 {
		for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
			output[i], output[j] = output[j], output[i]
		}
	}

	return output
}
