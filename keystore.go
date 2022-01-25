package keystore

import (
	"sync"
	"time"
)

// Simple in memeory KeyStore
//
// K is comparable as a requirment of map
// comparable is anything that satisfies != or ==
type KeyStore[K comparable, V any] struct {
	/*
		A RWMutex is a reader/writer mutual exclusion lock.
		The lock can be held by an arbitrary number of readers
		or a single writer. The zero value for a RWMutex is an unlocked mutex.

		A RWMutex must not be copied after first use.

		If a goroutine holds a RWMutex for reading and another goroutine
		might call Lock, no goroutine should expect to be able to acquire
		a read lock until the initial read lock is released. In particular,
		this prohibits recursive read locking. This is to ensure that the lock
		eventually becomes available; a blocked Lock call excludes new
		readers from acquiring the lock.
	*/
	mu sync.RWMutex

	//todo consider sync.Map
	//which is threadsafe, but optimised heavily towards reads
	store map[K]V
}

// Create keystore with K for key and V for Value
func NewKeyStore[K comparable, V any]() *KeyStore[K, V] {
	return &KeyStore[K, V]{
		store: make(map[K]V),
	}
}

// puts key/val pair into keystore using write lock
func (ks *KeyStore[K, V]) Put(key K, val V) {
	ks.mu.Lock()
	defer ks.mu.Unlock()
	ks.store[key] = val
}

// puts key/val and then deletes after d Duration
func (ks *KeyStore[K, V]) PutExpires(key K, val V, d time.Duration) {
	ks.Put(key, val)
	go func(key K, d time.Duration) {
		time.Sleep(d)
		ks.Delete(key)
	}(key, d)
}

// retrieves value using key
// bool is true if found
func (ks *KeyStore[K, V]) Get(key K) (V, bool) {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	//can't do this directly :(
	val, ok := ks.store[key]
	return val, ok
}

// Delete key from keystore
func (ks *KeyStore[K, V]) Delete(key K) {
	ks.mu.Lock()
	defer ks.mu.Unlock()
	delete(ks.store, key)
}
