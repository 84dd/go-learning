package main

import "fmt"

func main() {
	// map 会自动扩容
	map1 := map[string]int{"小明": 16, "小红": 18}
	fmt.Println(map1)

	// 利用make，创建容量为3的map，但底层会自动扩容
	map2 := make(map[string]int, 3)
	map2["小青1"] = 16
	map2["小青2"] = 17
	map2["小青3"] = 18
	map2["小青4"] = 19
	fmt.Println(map2)

	// 删除key
	delete(map2, "故意不存在的key")
	delete(map2, "小青4")
	fmt.Println(map2)

	// 判断key是否存在
	value, ok := map2["小青1"]
	fmt.Println(value, ok)

	// 遍历
	for k, v := range map2 {
		fmt.Println(k, " = ", v)
	}

}
