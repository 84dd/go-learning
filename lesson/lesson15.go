package lesson

import (
	"sync"
)

// 参考 乾航 分享的《Go 语言圣经》第九章《基于共享变量的并发》的内容
// 在你个人目录下创建 lesson15.go ，然后完成以下练习题：
// 已知 GO 语言中的 map 类型是非并发安全的，现在需要你设计一个新的 Cache 结构体来实现并发安全的类似 map 的数据结构
// •需要包含增/删/改/查;
// •Set() 方法用于将指定的键值对存储到缓存中;
// •Get() 方法用于根据键获取缓存中的值，并返回是否存在;
// •Delete() 方法用于删除指定键的缓存项;
// •多个 goroutine 并发访问缓存系统时不会出现竞态条件

// Cache 缓存结构体 自带同步锁
type Cache struct {
	mutex sync.RWMutex
	data  map[string]any
}

// NewCache 新建实例
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

// Set k、v 给 map 赋值
// 更新、插入
func (c *Cache) Set(key string, value any) {
	c.mutex.Lock()         // 获取写锁，写锁会阻塞读写操作
	defer c.mutex.Unlock() // 释放写锁
	c.data[key] = value    // 在获得锁后，进行set操作
}

// Get 通过 key 获取 value，并且通过 bool 类型判断 value 是否存在
// 查找
func (c *Cache) Get(key string) (any, bool) {
	c.mutex.RLock()         // 获取读锁，读锁不会阻塞读请求，会被写锁阻塞
	defer c.mutex.RUnlock() // 释放读锁
	val, ok := c.data[key]  // 获得读锁后，进行get操作
	return val, ok
}

// Delete 通过 key 删除键值对
// 删除
func (c *Cache) Delete(key string) {
	c.mutex.Lock()         // 获取写锁
	defer c.mutex.Unlock() // 释放写锁
	delete(c.data, key)    // 在获得锁后，操作删除
}
