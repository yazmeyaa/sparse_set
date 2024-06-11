package sparse_set

import "sync"

type SparseSet struct {
	size   int
	dense  []Id
	sparse []Id
	mx     sync.RWMutex
}

type Id = int

func NewSparseSet(maxSize uint32) *SparseSet {
	return &SparseSet{
		size:   0,
		sparse: make([]Id, maxSize),
		dense:  make([]Id, 0, maxSize),
	}
}

func (ss *SparseSet) Contains(x int) bool {
	ss.mx.RLock()
	defer ss.mx.RUnlock()

	return ss.contains(x)
}

func (ss *SparseSet) contains(x int) bool {
	return x < len(ss.sparse) && ss.sparse[x] < ss.size && ss.dense[ss.sparse[x]] == x
}

func (ss *SparseSet) Add(x int) {
	ss.mx.Lock()
	defer ss.mx.Unlock()

	if ss.contains(x) {
		return
	}
	ss.sparse[x] = ss.size
	ss.dense = append(ss.dense, x)
	ss.size++
}

func (ss *SparseSet) Remove(x int) {
	ss.mx.Lock()
	defer ss.mx.Unlock()

	if !ss.contains(x) {
		return
	}
	idx := ss.sparse[x]
	last := ss.dense[ss.size-1]
	ss.dense[idx] = last
	ss.sparse[last] = idx
	ss.size--
	ss.dense = ss.dense[:ss.size]
}

func (ss *SparseSet) Clear() {
	ss.mx.Lock()
	defer ss.mx.Unlock()

	ss.size = 0
}

func (ss *SparseSet) GetAll() []Id {
	ss.mx.RLock()
	defer ss.mx.RUnlock()

	elements := make([]int, ss.size)
	copy(elements, ss.dense[:ss.size])
	return elements
}

func (ss *SparseSet) Range(cb func(id Id)) {
	ss.mx.RLock()
	defer ss.mx.RUnlock()

	for _, id := range ss.dense[:ss.size] {
		cb(id)
	}
}
