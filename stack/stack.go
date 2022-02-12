package stack

import (
	"github.com/tonoy30/goutils/common"
	"sync"
)

type Stack struct {
	items  []common.Item
	rwLock sync.RWMutex
}

func New() *Stack {
	return &Stack{
		items:  []common.Item{},
		rwLock: sync.RWMutex{},
	}
}
func (s *Stack) Push(item common.Item) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.items = append(s.items, item)
}
func (s *Stack) Pop() common.Item {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	// Popping item from items slice.
	item := s.items[len(s.items)-1]
	//Adjusting the item's length accordingly
	s.items = s.items[0 : len(s.items)-1]
	return item
}
func (s *Stack) Size() int {
	// Acquire read lock
	s.rwLock.RLock()
	// defer operation of unlocking.
	defer s.rwLock.RUnlock()

	// Return length of items slice.
	return len(s.items)
}

// All - return all items present in stack
func (s *Stack) All() []common.Item {
	// Acquire read lock
	s.rwLock.RLock()

	// defer operation of unlocking.
	defer s.rwLock.RUnlock()

	// Return items slice to the caller.
	return s.items
}

// IsEmpty - Check is stack is empty or not.
func (s *Stack) IsEmpty() bool {
	// Acquire read lock
	s.rwLock.RLock()
	// defer operation of unlock.
	defer s.rwLock.RUnlock()

	return len(s.items) == 0
}

func Demo() []common.Item {
	stack := New()
	stack.Push(common.Item(1))
	return stack.All()
}
