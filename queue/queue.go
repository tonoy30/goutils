package queue

import (
	"github.com/tonoy30/goutils/common"
	"sync"
)

type Queue struct {
	items  []common.Item
	rwLock sync.RWMutex
}

func New() *Queue {
	return &Queue{
		items:  []common.Item{},
		rwLock: sync.RWMutex{},
	}
}

func (q *Queue) Enqueue(item common.Item) {
	q.rwLock.Lock()
	defer q.rwLock.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() common.Item {
	q.rwLock.Lock()
	defer q.rwLock.Unlock()
	// Popping item from items slice.
	item := q.items[0]
	//Adjusting the item's length accordingly
	q.items = q.items[1:]
	return item
}

func (q *Queue) Front() common.Item {
	q.rwLock.Lock()
	defer q.rwLock.Unlock()
	return q.items[0]
}
func (q *Queue) Size() int {
	// Acquire read lock
	q.rwLock.RLock()
	// defer operation of unlocking.
	defer q.rwLock.RUnlock()

	// Return length of items slice.
	return len(q.items)
}

// All - return all items present in stack
func (q *Queue) All() []common.Item {
	// Acquire read lock
	q.rwLock.RLock()
	// defer operation of unlocking.
	defer q.rwLock.RUnlock()
	// Return items slice to the caller.
	return q.items
}

// IsEmpty - Check is stack is empty or not.
func (q *Queue) IsEmpty() bool {
	// Acquire read lock
	q.rwLock.RLock()
	// defer operation of unlock.
	defer q.rwLock.RUnlock()

	return len(q.items) == 0
}

func Demo() []common.Item {
	q := New()
	q.Enqueue(common.Item(1))
	q.Enqueue(common.Item(2))
	q.Dequeue()
	return q.All()
}
