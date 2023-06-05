package lesson

import (
	"fmt"
	"testing"
)

func TestProductTemplate(t *testing.T) {
	jsonStr := `{
	 		"id": 1,
	 		"name": "Laptop",
	 		"price": 999.99
	 }`
	fmt.Println("\n结果：", ProductTemplate(jsonStr))
}
