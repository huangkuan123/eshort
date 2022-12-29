package elru

import (
	"time"
)

type Node struct {
	K     string
	V     interface{}
	Ltime int
	Num   int
}

func NewNode(k string, v interface{}) Node {
	return Node{K: k, V: v, Ltime: int(time.Now().Unix())}
}
