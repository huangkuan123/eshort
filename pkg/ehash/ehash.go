package ehash

import (
	"eshort/pkg/config"
	"eshort/pkg/murmur3"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

var chars string

func Setup() {
	chars = config.GetString("eshort.key")
}

func HashKey(id uint64) string {
	//防猜测遍历
	seed := rand.Uint32()
	str := strconv.FormatUint(id, 10) + strconv.FormatInt(rand.Int63(), 10)
	sum := murmur3.Sum32WithSeed([]byte(str), seed)
	short := encode(int64(sum))
	return short
}

func Decode(str string) int64 {
	var num int64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(chars, str[i])
		num += int64(math.Pow(62, float64(n-i-1)) * float64(pos))
	}
	return num
}

func encode(num int64) string {
	bytes := []byte{}
	for num > 0 {
		bytes = append(bytes, chars[num%62])
		num = num / 62
	}
	reverse(bytes)
	return string(bytes)
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}
