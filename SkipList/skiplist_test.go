package skiplist

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestBasicOperations(t *testing.T) {
	s := &Skiplist{
		head: &node{
			nexts: make([]*node, 0),
		},
	}

	s.Put(1, 10)
	s.Put(2, 20)
	s.Put(3, 30)

	if val, ok := s.Get(2); !ok || val != 20 {
		t.Errorf("Get(2) = %v, %v; want 20, true", val, ok)
	}

	s.Put(2, 25)
	if val, ok := s.Get(2); !ok || val != 25 {
		t.Errorf("Get(2) after Put(2, 25) = %v, %v; want 25, true", val, ok)
	}

	s.Del(2)
	if _, ok := s.Get(2); ok {
		t.Errorf("Get(2) after Del(2) returned true, want false")
	}
}

func TestRangeQuery(t *testing.T) {
	s := &Skiplist{
		head: &node{
			nexts: make([]*node, 0),
		},
	}
	for i := 0; i < 100; i++ {
		s.Put(i, i*10)
	}

	range_result := s.Range(20, 29)
	if len(range_result) != 10 {
		t.Errorf("Range(20, 29) returned %d elements, want 10", len(range_result))
	}
	for i, kv := range range_result {
		if kv[0] != 20+i || kv[1] != (20+i)*10 {
			t.Errorf("Range(20, 29)[%d] = %v, want [%d, %d]", i, kv, 20+i, (20+i)*10)
		}
	}
}

func TestCeilingAndFloor(t *testing.T) {
	s := &Skiplist{
		head: &node{
			nexts: make([]*node, 0),
		},
	}
	s.Put(10, 100)
	s.Put(20, 200)
	s.Put(30, 300)

	if kv, ok := s.Ceiling(15); !ok || kv != [2]int{20, 200} {
		t.Errorf("Ceiling(15) = %v, %v; want [20, 200], true", kv, ok)
	}

	if kv, ok := s.Floor(25); !ok || kv != [2]int{20, 200} {
		t.Errorf("Floor(25) = %v, %v; want [20, 200], true", kv, ok)
	}
}

func TestPerformance(t *testing.T) {
	s := &Skiplist{
		head: &node{
			nexts: make([]*node, 0),
		},
	}
	n := 100000
	keys := make([]int, n)

	// 插入性能
	start := time.Now()
	for i := 0; i < n; i++ {
		key := rand.Int()
		s.Put(key, i)
		keys[i] = key
	}
	insertDuration := time.Since(start)
	t.Logf("Insert %d elements: %v", n, insertDuration)

	// 查询性能
	start = time.Now()
	for _, key := range keys {
		s.Get(key)
	}
	getDuration := time.Since(start)
	t.Logf("Get %d elements: %v", n, getDuration)

	// 删除性能
	start = time.Now()
	for _, key := range keys {
		s.Del(key)
	}
	deleteDuration := time.Since(start)
	t.Logf("Delete %d elements: %v", n, deleteDuration)
}

func TestCorrectness(t *testing.T) {
	s := &Skiplist{
		head: &node{
			nexts: make([]*node, 0),
		},
	}
	n := 10000
	keys := make([]int, n)

	for i := 0; i < n; i++ {
		key := rand.Int() % (n * 10)
		s.Put(key, i)
		keys[i] = key
	}

	sort.Ints(keys)

	// 测试Range
	start, end := keys[n/4], keys[3*n/4]
	range_result := s.Range(start, end)
	for i := 1; i < len(range_result); i++ {
		if range_result[i][0] <= range_result[i-1][0] {
			t.Errorf("Range result not sorted: %v <= %v", range_result[i][0], range_result[i-1][0])
		}
	}

	// 测试Ceiling和Floor
	for i := 0; i < 100; i++ {
		target := rand.Int() % (n * 10)
		ceiling, cok := s.Ceiling(target)
		floor, fok := s.Floor(target)

		if cok && ceiling[0] < target {
			t.Errorf("Ceiling(%d) = %v, should be >= %d", target, ceiling, target)
		}
		if fok && floor[0] > target {
			t.Errorf("Floor(%d) = %v, should be <= %d", target, floor, target)
		}
		if cok && fok && ceiling[0] < floor[0] {
			t.Errorf("Ceiling(%d) = %v, Floor(%d) = %v, ceiling should be >= floor", target, ceiling, target, floor)
		}
	}
}
