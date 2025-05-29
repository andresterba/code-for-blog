package main

import (
	"hash/fnv"
	"log"
	"sync"
)

const (
	baseCapicity = 16
)

type Entry struct {
	key   string
	value any
	next  *Entry
}

type HashMap struct {
	buckets  []*Entry
	size     int
	capacity int
	mutex    sync.RWMutex
}

func NewHashMap() *HashMap {
	return &HashMap{
		buckets:  make([]*Entry, baseCapicity),
		capacity: baseCapicity,
		size:     0,
	}
}

func NewHashMapWithCapacity(capacity int) *HashMap {
	return &HashMap{
		buckets:  make([]*Entry, capacity),
		capacity: capacity,
		size:     0,
	}
}

func hashFunction(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (hm *HashMap) resize() {
	oldBuckets := hm.buckets
	hm.capacity *= 2
	hm.buckets = make([]*Entry, hm.capacity)
	hm.size = 0 // Reset size, will be recalculated during rehashing

	for _, entry := range oldBuckets {
		for e := entry; e != nil; e = e.next {
			log.Println("Rehashing key:", e.key, "with value:", e.value)
			hm.Insert(e.key, e.value) // Reinsert entries into the new buckets
		}
	}
}

func (hm *HashMap) Insert(key string, value any) {
	if hm.size == hm.capacity {
		log.Println("HashMap is full, resizing...")
		hm.resize()
	}

	hm.mutex.Lock()
	defer hm.mutex.Unlock()

	hashResult := hashFunction(key)
	log.Println("Hash result for key:", key, "is", hashResult)
	index := int(hashFunction(key)) % hm.capacity
	log.Println("Index for key:", key, "is", index)

	log.Println("Putting key:", key, "with value:", value, "at index:", index)
	head := hm.buckets[index]

	// Update existing key
	for e := head; e != nil; e = e.next {
		log.Println("Checking entry:", e.key)
		if e.key == key {
			e.value = value
			return
		}
	}

	// The magic part happens when we insert a new entry into the bucket.
	// We will set its next pointer to the current head of the bucket.
	// By this the new entry will be the first entry in the linked list,
	// and the previous entries will be accessible via the next pointers.
	newEntry := &Entry{
		key:   key,
		value: value,
		next:  head,
	}
	hm.buckets[index] = newEntry
	hm.size++
}

func (hm *HashMap) Get(key string) (any, bool) {
	hm.mutex.RLock()
	defer hm.mutex.RUnlock()

	hashResult := hashFunction(key)
	log.Println("Hash result for key:", key, "is", hashResult)
	index := int(hashFunction(key)) % hm.capacity
	log.Println("Getting key:", key, "at index:", index)
	head := hm.buckets[index]

	for e := head; e != nil; e = e.next {
		log.Println("Checking entry:", e.key)
		if e.key == key {
			return e.value, true
		}
	}

	return nil, false
}

func (hm *HashMap) Delete(key string) bool {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()

	index := int(hashFunction(key)) % hm.capacity
	head := hm.buckets[index]
	var prev *Entry

	for e := head; e != nil; e = e.next {
		if e.key == key {
			if prev == nil {
				hm.buckets[index] = e.next
			} else {
				prev.next = e.next
			}
			hm.size--
			return true
		}
		prev = e
	}

	return false
}

func (hm *HashMap) GetAll() []Entry {
	hm.mutex.RLock()
	defer hm.mutex.RUnlock()

	var entries []Entry
	for _, bucket := range hm.buckets {
		for e := bucket; e != nil; e = e.next {
			entries = append(entries, *e)
		}
	}

	return entries
}
