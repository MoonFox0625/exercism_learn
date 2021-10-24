package robotname

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

var nameBox = map[string]bool{}
var maxGetNameTime = 5 * time.Second

type Robot struct {
	name string
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randNumber(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += strconv.Itoa(rand.Intn(10))
	}
	return str
}

func (r *Robot) Name() (string, error) {
	// rand.Seed(time.Now().UnixNano())

	if r.name == "" {
		start := time.Now()
		for time.Now().Unix() < start.Add(maxGetNameTime).Unix() {
			var name = randSeq(2) + randNumber(3)
			if nameBox[name] == true {
				continue
			} else {
				r.name = name
				nameBox[r.name] = true
				return r.name, nil
			}
		}
		return "", errors.New("running overtime")
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	// nameBox[r.name] = false
	r.name = ""
}
