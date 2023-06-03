package lesson

import (
	"sync"
	"testing"
)

func TestCache(t *testing.T) {
	c := NewCache()

	var wg sync.WaitGroup //创建一个等待组
	wg.Add(4)             //等待组计数器+n

	for i := 0; i < 2; i++ {
		go func() {
			c.Set("k1", "v1")
			c.Delete("k1")
			c.Set(string("k2"), "v2")
			c.Set(string("k2"), "vv2") //更新
			defer wg.Done()            //计数器减1

		}()
	}
	for j := 0; j < 2; j++ {
		go func() {
			defer wg.Done()
			c.Set("k3", "v3")
			c.Set(string("k3"), "vv3") //更新
			c.Delete("k3")             //删除

		}()
	}
	wg.Wait()

	c.Set("k-string", "kkvv")
	val1, _ := c.Get("k-string")
	println("any类型的输出：", val1)
	println("any类型转string", val1.(string))

	c.Set("k-int", 22)
	val2, _ := c.Get("k-int")
	println("any类型的输出：", val2)
	println("any类型转string", val2.(int))
}
