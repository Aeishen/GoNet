package main

/*
Gob 是 Go 自己的以二进制形式序列化和反序列化程序数据的格式；可以在 encoding 包中找到。这种格式的数据简称为 Gob （即 Go binary 的缩写）
Gob 特定地用于纯 Go 的环境中，例如，两个用 Go 写的服务之间的通信。这样的话服务可以被实现得更加高效和优化。
Gob 的首选格式是二进制，而不是像 JSON 和 XML 那样的文本格式。
Gob 在编码和解码过程中用到了 Go 的反射。
Gob 文件或流是完全自描述的：里面包含的所有类型都有一个对应的描述，并且总是可以用 Go 解码，而不需要了解文件的内容。
只有可导出的字段会被编码，零值会被忽略。在解码结构体的时候，只有同时匹配名称和可兼容类型的字段才会被解码。当源数据类型增加新字段后，
Gob 解码客户端仍然可以以这种方式正常工作：解码客户端会继续识别以前存在的字段。并且还提供了很大的灵活性，比如在发送者看来，整数被编码成没有固定长度的可变长度，而忽略具体的 Go 类型。
*/

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"sync"
)

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
	file *os.File
}

type record struct {
	Key, URL string
}

func NewURLStore(filename string) *URLStore {
	s := &URLStore{urls: make(map[string]string)}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening URLStore:", err)
	}
	s.file = f
	if err := s.load(); err != nil {
		log.Println("Error loading URLStore:", err)
	}
	return s
}

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[key]; present {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if ok := s.Set(key, url); ok {
			if err := s.save(key, url); err != nil {
				log.Println("Error saving to URLStore:", err)
			}
			return key
		}
	}
}

func (s *URLStore) load() error {

	//从文件的开头寻找、读取并解码每一条记录
	if _, err := s.file.Seek(0, 0); err != nil {
		return err
	}
	d := gob.NewDecoder(s.file)
	var err error
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			s.Set(r.Key, r.URL)
		}
	}

	// io.EOF 错误是指解码文件结束
	if err == io.EOF {
		return nil
	}
	return err
}

func (s *URLStore) save(key, url string) error {
	e := gob.NewEncoder(s.file)
	return e.Encode(record{key, url})
}