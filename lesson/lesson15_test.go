package lesson

import (
	"fmt"
	"sync"
	"testing"
)

func TestCache(t *testing.T) {
	cache := NewCache()

	fmt.Println("单线程下测试cache的Set,Get,Delete功能")
	cache.Set("k1", 11)
	cache.Set("k1", 22)
	fmt.Println(cache.Get("k1"))
	cache.Delete("k1")
	fmt.Println(cache.Get("k1"))

	fmt.Println("多线程测试")
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(3)

		go func() {
			cache.Set("k2", i)
			wg.Done()
		}()
		go func() {
			cache.Get("k2")
			wg.Done()
		}()
		go func() {
			cache.Delete("k2")
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			cache.Set("k3", 33)
			cache.Get("k3")
			cache.Delete("k3")
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cache.data)

	cache.Set("k-string", "kkvv")
	val1, _ := cache.Get("k-string")
	println("any类型的输出：", val1)
	println("any类型转string", val1.(string))

	cache.Set("k-int", 22)
	val2, _ := cache.Get("k-int")
	println("any类型的输出：", val2)
	println("any类型转string", val2.(int))
}
