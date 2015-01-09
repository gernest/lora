package number2id

import (
	"bytes"
	"math"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	floor   = math.Floor
	log     = math.Log
	mod     = math.Mod
	pow     = math.Pow
	indexOf = strings.Index
)

func Encode(n int64) string {
	result := bytes.NewBuffer([]byte{})
	f := float64(n)
	al := float64(len(alphabet))
	for i := floor(log(f) / log(al)); i >= 0; i-- {
		idx := int(mod(floor(f/bpow(al, i)), al))
		result.WriteString(alphabet[idx : idx+1])
	}

	return reverseString(result.String())
}

func Decode(id string) int64 {
	str := reverseString(id)
	result := int64(0)
	end := len(str) - 1
	al := float64(len(alphabet))
	for i := 0; i <= end; i++ {
		result += int64(float64(indexOf(alphabet, str[i:i+1])) * bpow(al, float64(end-i)))
	}

	return result
}

func bpow(a float64, b float64) float64 {
	return floor(pow(a, b))
}

func reverseString(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
