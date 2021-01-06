package homework

import (
	"sync"
	"time"
)

// Rolling 滑动窗口的简单实现
type Rolling struct {
	buckets        []*bucket
	last           time.Time     // current 上一次变动时的时间
	current        int           // 当前位置
	size           int           // 桶的数量，默认为 10
	bucketDuration time.Duration // 桶的时间间隔, 默认 1s
	mu             sync.RWMutex  // 使用读写锁实现，后续如果有更高要求可以使用 atomic 包
}

// bucket 用 struct 实现，后面方便扩展
type bucket struct {
	val float64
}

// Reset Reset
func (b *bucket) Reset() {
	b.val = 0
}

// RollingOption RollingOption
type RollingOption func(opt *Rolling)

// WithBucketDuration WithBucketDuration
func WithBucketDuration(t time.Duration) RollingOption {
	return func(opt *Rolling) {
		opt.bucketDuration = t
	}
}

// NewRolling NewRolling
func NewRolling(opts ...RollingOption) *Rolling {
	rolling := &Rolling{
		size:           10,
		bucketDuration: time.Second,
		last:           time.Now(),
	}

	for _, opt := range opts {
		opt(rolling)
	}

	rolling.buckets = make([]*bucket, rolling.size)
	for i := range rolling.buckets {
		rolling.buckets[i] = &bucket{}
	}
	return rolling
}

// currentBucket 获取当前的 currentBucket
func (r *Rolling) currentBucket() *bucket {
	old := r.current
	// 计算需要往前走多少步
	s := int(time.Since(r.last) / r.bucketDuration)
	if s > 0 {
		r.last = time.Now()
	}

	// 计算新的 current
	r.current = (old + s) % r.size

	// 避免 s 过大时空转
	if s > r.size {
		s = r.size
	}

	// 清空前面走过的路
	for i := 1; i <= s; i++ {
		r.buckets[(old+i)%r.size].Reset()
	}
	return r.buckets[r.current]
}

// Add 添加一个自定义数值
func (r *Rolling) Add(val float64) {
	if val == 0 {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.currentBucket().val += val
}

// Sum 数据统计
func (r *Rolling) Sum() float64 {
	var sum float64

	r.mu.RLock()
	defer r.mu.RUnlock()
	old := r.current
	s := int(time.Since(r.last) / r.bucketDuration)
	// 计算新的 current
	n := (old + s) % r.size

	// 求和
	for i := 0; i < r.size-s; i++ {
		sum += r.buckets[(n+i+1)%r.size].val
	}
	return sum
}
