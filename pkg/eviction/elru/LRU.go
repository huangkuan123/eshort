package elru

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var ELRU LRU

type LRU struct {
	cap               int //容量
	count             int //实际占用
	evictionThreshold int //开始逐出
	data              map[string]Node
	pool              pool
}

type pool struct {
	data []Node
}

func NewLRU(cap int) LRU {
	lru := LRU{cap: cap, evictionThreshold: cap / 2, pool: pool{}, data: make(map[string]Node)}
	ELRU = lru
	return lru
}

func (a *LRU) Put(n Node) (bool, error) {
	if a.cap == 0 { //
		return false, errors.New("未初始化")
	}
	a.data[n.K] = n
	if a.count == a.cap {
		//立即执行逐出
		a.count += 1
		return true, nil
	}
	a.count += 1
	if a.count > a.evictionThreshold {
		a.joinPool()
	}
	if a.count >= a.cap {
		a.eviction()
	}
	return true, nil
}

// 溢出
func (a *LRU) eviction() {
	num := 3
	lpool := len(a.pool.data)
	fmt.Println(lpool, a.pool)
	es := a.pool.data[0:num]
	fmt.Println(es)
	a.pool.data = a.pool.data[num : lpool-1]
	fmt.Println(a.pool.data)
	a.count -= num
	for _, node := range es {
		delete(a.data, node.K)
	}
}

func (a *LRU) Get(key string) (interface{}, bool) {
	if len(key) == 0 {
		return "", false
	}
	node, ok := a.data[key]
	if !ok {
		return "", false
	}
	node.Ltime = int(time.Now().Unix())
	node.Num += 1
	a.data[key] = node
	return node.V, true
}

func (a *LRU) joinPool() {
	randkeys := a.randkey()
	f := false
	for _, node := range randkeys {
		if node.Ltime <= a.pool.data[0].Ltime {
			f = true
			a.pool.data = append(a.pool.data, node)
		}
	}
	if f {
		sort.SliceStable(a.data, func(i, j int) bool {
			return a.pool.data[i].Ltime < a.pool.data[j].Ltime
		})
	}
	ldata := len(a.pool.data)
	poolLen := 16
	if ldata >= poolLen {
		//清理淘汰池
		a.pool.data = a.pool.data[0:15]
	}
}

func (a *LRU) randkey() []Node {
	rs := [5]int{rand.Intn(a.count), rand.Intn(a.count), rand.Intn(a.count), rand.Intn(a.count), rand.Intn(a.count)}
	jp := []Node{}
	for _, r := range rs {
		t := 0
		for _, node := range a.data {
			if t == r {
				jp = append(jp, node)
				break
			}
			t += 1
		}
	}
	return jp
}
