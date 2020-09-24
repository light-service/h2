package unblock

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ErrExpired = errors.New("mutex expired")

	redisPrefix = ""
)

func SetRedisPrefix(prefix string) {
	redisPrefix = prefix
}

type RedisMutex struct {
	client    *redis.Client
	name      string
	expiry    time.Duration
	startTime time.Time
	isOwner   bool
}

func NewRedisMutex(client *redis.Client, name string, expiry time.Duration) *RedisMutex {
	if client == nil {
		panic("client is nil")
	}

	return &RedisMutex{
		client:  client,
		name:    name,
		expiry:  expiry,
		isOwner: false,
	}
}

func (m *RedisMutex) TryLock() error {
	if ok, _ := m.client.SetNX(context.TODO(), m.key(), 1, m.expiry).Result(); !ok {
		return ErrAlreadyLocked
	}

	m.startTime = time.Now()
	m.isOwner = true
	return nil

}

func (m *RedisMutex) TryUnlock() error {
	if !m.isOwner {
		return ErrNotOwner
	}
	if time.Since(m.startTime) >= m.expiry {
		return ErrExpired
	}

	if err := m.client.Del(context.TODO(), m.key()).Err(); err != nil {
		return err
	}

	m.isOwner = false
	return nil
}

func (m RedisMutex) key() string {
	if redisPrefix != "" {
		return redisPrefix + m.name
	} else {
		return "redis_lock:" + m.name
	}
}
