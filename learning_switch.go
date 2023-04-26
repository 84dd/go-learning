package main

import "fmt"

func main() {
	switch getDay() {
	case 1:
		fmt.Println("星期一")
		// 不加 break，不会造成case穿透
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	}

	switch getDay() {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	// 不加 break，不会造成case穿透
	case 6:
		fallthrough // 显式case穿透
	case 7:
		fallthrough
	default:
		fmt.Println("周末")
	}

	switch getDay() {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
		// 不加 break，不会造成case穿透
	default:
		fmt.Println("周末")
	}
}

func getDay() int {
	return 1
}

func getDay2() int {
	return 2
}
