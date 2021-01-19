package queue

import "sync"

type QueuePool struct {
	pool *Pool
	lock sync.Mutex
}

var (
	stackOnce sync.Once
	stackPool *QueuePool
)

// Single 单例初始化
func Single() *QueuePool {
	stackOnce.Do(func() {
		stackPool = NewQueue()
	})
	return stackPool
}

// NewQueue 实例化队列
func NewQueue() *QueuePool {
	return &QueuePool{
		lock: sync.Mutex{},
		pool: new(Pool),
	}
}

// Push
func (s *QueuePool) Push(x interface{}) {
	s.lock.Lock()
	s.pool.Push(x)
	s.lock.Unlock()
}

// Pop
func (s *QueuePool) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.pool.Len() > 0 {
		return s.pool.Pop()
	}
	return nil
}

// Len
func (s *QueuePool) Len() int {
	return s.pool.Len()
}

type Pool []interface{}

func (p Pool) Len() int { return len(p) }

func (p Pool) Less(i, j int) bool { return true }

func (p Pool) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p *Pool) Push(x interface{}) { *p = append(*p, x) }

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}
