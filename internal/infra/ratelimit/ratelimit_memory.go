package ratelimit

import (
	"sync"
	"time"
)

type memoryBucket struct {
	count        int
	expiresAt    time.Time
	blockedUntil time.Time
}

type MemoryLimiter struct {
	mu       sync.Mutex
	buckets  map[string]*memoryBucket
	limit    int
	duration time.Duration
	blockFor time.Duration
}

func NewMemoryLimiter(limit int, duration time.Duration, blockFor time.Duration) *MemoryLimiter {
	return &MemoryLimiter{
		buckets:  make(map[string]*memoryBucket),
		limit:    limit,
		duration: duration,
		blockFor: blockFor,
	}
}

func (m *MemoryLimiter) Rate(ip string, token string) (bool, error) {
	key := ip
	if token != "" {
		key = "token:" + token
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	b, exists := m.buckets[key]

	if !exists || now.After(b.expiresAt) {
		m.buckets[key] = &memoryBucket{
			count:        1,
			expiresAt:    now.Add(m.duration),
			blockedUntil: time.Time{},
		}
		return true, nil
	}

	if now.Before(b.blockedUntil) {
		return false, nil
	}

	if b.count >= m.limit {
		b.blockedUntil = now.Add(m.blockFor)
		return false, nil
	}

	b.count++
	return true, nil
}
