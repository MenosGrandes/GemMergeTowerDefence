package pool

import (
	"example/hello/src/constants"
	"sync"
)

const (
	PoolSize = 100
)

type FixedPool[T any] struct {
	mu      sync.Mutex
	objects [PoolSize]T
	top     constants.ID
}

func NewFixedPool[T any](create func(i constants.ID) T) *FixedPool[T] {
	p := &FixedPool[T]{
		top: PoolSize,
	}
	for i := range constants.ID(PoolSize) {
		p.objects[i] = create(i)
	}
	return p
}
func (p *FixedPool[T]) Put(obj *T) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.top == constants.ID(len(p.objects)) {
		return
	}
	p.objects[p.top] = *obj
	p.top++
}
func (p *FixedPool[T]) Get() *T {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.top == 0 {
		return nil
	}
	p.top--
	return &p.objects[p.top]
}
