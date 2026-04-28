package utils

import (
	"sync"
)

var (
	GlobalCache = make(map[string][]byte)
	cacheMutex  sync.RWMutex
)

// GetCache retrieves JSON data from the memory cache based on the route key.
func GetCache(key string) []byte {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	return GlobalCache[key]
}

// SetCache stores JSON data into the memory cache.
func SetCache(key string, data []byte) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	
	// Create a copy of the data to avoid slice mutation issues
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)
	
	GlobalCache[key] = dataCopy
}
