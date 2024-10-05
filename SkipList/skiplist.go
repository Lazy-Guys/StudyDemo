package skiplist

import (
	"math/rand"
)

type node struct {
	nexts []*node
	key   int
	val   int
}

type Skiplist struct {
	head *node
}

func (s *Skiplist) Get(key int) (int, bool) {
	if _node := s.search(key); _node != nil {
		return _node.val, true
	}
	return -1, false
}

func (s *Skiplist) search(key int) *node {
	n := s.head
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		for n.nexts[level] != nil && n.nexts[level].key < key {
			n = n.nexts[level]
		}

		// 若找到目标值，则返回
		if n.nexts[level] != nil && n.nexts[level].key == key {
			return n.nexts[level]
		}

		// 若未找到目标值，则继续向下遍历
	}
	return nil
}

func roll() int {
	var level int
	for rand.Float32() < 0.5 && level < 32 {
		level++
	}
	return level
}

func (s *Skiplist) Put(key int, val int) {
	if _node := s.search(key); _node != nil {
		_node.val = val
		return
	}
	level := roll()
	for len(s.head.nexts)-1 < level {
		s.head.nexts = append(s.head.nexts, nil)
	}
	newNode := &node{
		key:   key,
		val:   val,
		nexts: make([]*node, level+1),
	}
	n := s.head
	for l := level; l >= 0; l-- {
		for n.nexts[l] != nil && n.nexts[l].key < key {
			n = n.nexts[l]
		}
		newNode.nexts[l] = n.nexts[l]
		n.nexts[l] = newNode
	}
}

func (s *Skiplist) Del(key int) {
	if _node := s.search(key); _node == nil {
		return
	}
	n := s.head
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		for n.nexts[level] != nil && n.nexts[level].key < key {
			n = n.nexts[level]
		}
		if n.nexts[level] == nil || n.nexts[level].key > key {
			continue
		}
		n.nexts[level] = n.nexts[level].nexts[level]
	}

	var dif int
	for level := len(s.head.nexts) - 1; level > 0 && s.head.nexts[level] == nil; level-- {
		dif++
	}
	s.head.nexts = s.head.nexts[:len(s.head.nexts)-dif]
}

func (s *Skiplist) Range(start, end int) [][2]int {
	ceilNode := s.ceiling(start)
	if ceilNode == nil {
		return [][2]int{}
	}

	var res [][2]int
	for n := ceilNode; n != nil && n.key <= end; n = n.nexts[0] {
		res = append(res, [2]int{n.key, n.val})
	}
	return res
}

func (s *Skiplist) ceiling(start int) *node {
	n := s.head
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		for n.nexts[level] != nil && n.nexts[level].key < start {
			n = n.nexts[level]
		}
		if n.nexts[level] != nil && n.nexts[level].key == start {
			return n.nexts[level]
		}
	}
	return n.nexts[0]
}

func (s *Skiplist) Ceiling(target int) ([2]int, bool) {
	if ceilNode := s.ceiling(target); ceilNode != nil {
		return [2]int{ceilNode.key, ceilNode.val}, true
	}
	return [2]int{}, false
}

func (s *Skiplist) Floor(target int) ([2]int, bool) {
	if floorNode := s.floor(target); floorNode != nil {
		return [2]int{floorNode.key, floorNode.val}, true
	}
	return [2]int{}, false
}

func (s *Skiplist) floor(target int) *node {
	n := s.head
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		for n.nexts[level] != nil && n.nexts[level].key < target {
			n = n.nexts[level]
		}
		if n.nexts[level] != nil && n.nexts[level].key == target {
			return n.nexts[level]
		}
	}
	return n
}
