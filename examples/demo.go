package main

import (
	"fmt"
	"time"

	goif "github.com/shengyanli1982/go-if"
)

// 辅助函数：模拟可能出错的操作
// Helper function: simulate operation that might fail
func someOperation() error {
	// 这里模拟一个始终成功的操作
	// Simulate an always-successful operation
	return nil
}

func main() {
	// 1. 基础类型使用示例
	// Basic type examples
	name := "World"
	greeting := goif.If(len(name) > 0,
		fmt.Sprintf("Hello, %s!", name),
		"Hello, Guest!")
	fmt.Println(greeting) // 输出: Hello, World!

	// 2. 数值计算示例
	// Numeric calculation example
	score := 75
	result := goif.If(score >= 60,
		"Pass",
		"Fail")
	fmt.Printf("Score: %d, Result: %s\n", score, result) // 输出: Score: 75, Result: Pass

	// 3. 时间判断示例
	// Time condition example
	currentHour := time.Now().Hour()
	greeting = goif.If(currentHour < 12,
		"Good morning!",
		"Good afternoon!")
	fmt.Println(greeting)

	// 4. 切片操作示例
	// Slice operation example
	numbers := []int{1, 2, 3}
	result = goif.If(len(numbers) > 0,
		fmt.Sprintf("First number is: %d", numbers[0]),
		"Slice is empty")
	fmt.Println(result) // 输出: First number is: 1

	// 5. 错误处理示例
	// Error handling example
	err := someOperation()
	message := goif.If(err == nil,
		"Operation successful",
		fmt.Sprintf("Operation failed: %v", err))
	fmt.Println(message)

	// 6. 结构体示例
	// Struct example
	type User struct {
		Name    string
		IsAdmin bool
	}
	user := User{Name: "Alice", IsAdmin: true}
	access := goif.If(user.IsAdmin,
		"Full access granted",
		"Limited access")
	fmt.Printf("User %s: %s\n", user.Name, access) // 输出: User Alice: Full access granted

	// 7. 嵌套使用示例
	// Nested usage example
	age := 25
	status := goif.If(age >= 18,
		goif.If(age <= 60,
			"Adult",
			"Senior"),
		"Minor")
	fmt.Printf("Age %d: %s\n", age, status) // 输出: Age 25: Adult

	// 8. 函数返回值示例
	// Function return value example
	calculate := func(x int) int {
		return goif.If(x > 0,
			x*2,
			-x)
	}
	fmt.Printf("Calculate(5): %d\n", calculate(5))   // 输出: Calculate(5): 10
	fmt.Printf("Calculate(-3): %d\n", calculate(-3)) // 输出: Calculate(-3): 3
}
