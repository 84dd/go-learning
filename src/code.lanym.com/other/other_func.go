package other

import "fmt"

// FinelyPrint 优雅地打印
//
// param: name
// param: age
func FinelyPrint(name string, age int) {
	fmt.Printf("%v今年%v岁\n", name, age)
}
