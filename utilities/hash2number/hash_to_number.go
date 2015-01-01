package hash2number

import (
	"hash/fnv"
)

func Hash2number(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
