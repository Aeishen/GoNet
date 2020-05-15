package main

import (
	"fmt"
	"sync"
)

type URLStore struct {
	urls map[string]string
	mu *sync.RWMutex  // 一个 RWMutex 有两个锁： 一个用于读取，一个用于写入。多个客户端可以同时获得读取锁，但是只能有一个客户端能够获得写入锁（禁用所有读取器），从而有效的序列化更新，使他们连续的工作。
}

func NewURLStore() *URLStore {
	return &URLStore{make(map[string]string),new(sync.RWMutex)}
}

// 重定向请求
func (s *URLStore)Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

//
func (s *URLStore)Set(key,url string) bool{
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, isFound := s.urls[key]; isFound{
		return false
	}
	s.urls[key] = url
	fmt.Printf("URLStore : %#v\n", s.urls)
	return true
}

//
func (s *URLStore)Put(url string)string{
	for{
		key := genKey(s.Count())
		if s.Set(key, url) {
			return key
		}
	}
}

//
func (s *URLStore)Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}
