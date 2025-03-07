package main

import (
	cacheServer "memCache/cache-server"
	"time"
)

func main() {
	cache := cacheServer.NewMemCache()
	cache.SetMaxMemory("100MB")

	cache.Set("int", 1, time.Second)
	cache.Set("bool", false, time.Second)
	cache.Set("data", map[string]interface{}{"a": 1}, time.Second)
	cache.Set("int", 1)
	cache.Set("bool", false)
	cache.Set("data", map[string]interface{}{"a": 1})

	cache.Get("int")
	cache.Get("bool")
	//cache.Del("int")
	//cache.Flush()
	//cache.Keys()
}
