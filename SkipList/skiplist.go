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

	// 随机生成层数
	level := roll()

	// 若层数大于当前层数，则扩充层数
	for len(s.head.nexts)-1 < level {
		s.head.nexts = append(s.head.nexts, nil)
	}

	// 创建新节点
	newNode := &node{
		key:   key,
		val:   val,
		nexts: make([]*node, level+1),
	}

	// 将新节点插入到跳表中
	n := s.head
	for l := level; l >= 0; l-- {
		for n.nexts[l] != nil && n.nexts[l].key < key {
			n = n.nexts[l]
		}

		// 当前节点n即为符合条件的节点
		// 将新节点插入到当前节点之后
		newNode.nexts[l] = n.nexts[l]
		n.nexts[l] = newNode
	}
}

func (s *Skiplist) Del(key int) {
	// 若未找到目标值，则无需删除，直接返回
	if _node := s.search(key); _node == nil {
		return
	}
	n := s.head
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		for n.nexts[level] != nil && n.nexts[level].key < key {
			n = n.nexts[level]
		}

		// 若未找到目标值，则无需删除，直接返回
		if n.nexts[level] == nil || n.nexts[level].key > key {
			continue
		}

		// 删除节点
		n.nexts[level] = n.nexts[level].nexts[level]
	}

	// 删除空层
	var dif int
	// 这个for循环的目的是找到所有头节点后直接指向nil的层数
	for level := len(s.head.nexts) - 1; level > 0 && s.head.nexts[level] == nil; level-- {
		dif++
	}
	// 将头结点中多余的层数删除
	// 这里有个疑问，也就是只删除了头节点中多余的层数，但是后继节点没有删除这些层
	// 实际上，无论是搜索、插入还是删除，所有的遍历都是由头结点的层数来指定的
	// 后继结点的这些层数在遍历过程中不会被访问到
	// 因此，这些层数并不会影响到跳表的正确性
	// 但是，这些层数占用的空间没有被释放
	// 随着跳表中元素的增多，这部分内存可能会影响跳表的性能
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

	// 从最高层开始遍历
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		// 遍历当前层，找到第一个大于等于start的节点
		for n.nexts[level] != nil && n.nexts[level].key < start {
			n = n.nexts[level]
		}

		// 若找到目标值，则返回
		if n.nexts[level] != nil && n.nexts[level].key == start {
			return n.nexts[level]
		}
	}

	// 若未找到目标值，则返回最低层的第一个节点
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

	// 从最高层开始遍历
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		// 遍历当前层，找到第一个大于等于target的节点
		for n.nexts[level] != nil && n.nexts[level].key < target {
			n = n.nexts[level]
		}

		// 若找到目标值，则返回
		if n.nexts[level] != nil && n.nexts[level].key == target {
			return n.nexts[level]
		}
	}

	// 目前n指向的就是小于target的最大节点，将其返回即可
	return n
}
