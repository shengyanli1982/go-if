package goif

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	// Basic types testing
	// 基础类型测试
	t.Run("string type", func(t *testing.T) {
		assert.Equal(t, "yes", If(true, "yes", "no"))
		assert.Equal(t, "no", If(false, "yes", "no"))
		// Test empty strings
		// 测试空字符串
		assert.Equal(t, "", If(true, "", "non-empty"))
		assert.Equal(t, "", If(false, "non-empty", ""))
	})

	t.Run("integer type", func(t *testing.T) {
		// Basic integer tests
		// 基础整数测试
		assert.Equal(t, 1, If(true, 1, 0))
		assert.Equal(t, 0, If(false, 1, 0))

		// Boundary values for int
		// 整数边界值测试
		assert.Equal(t, math.MaxInt64, If(true, math.MaxInt64, 0))
		assert.Equal(t, math.MinInt64, If(false, 0, math.MinInt64))

		// Test with negative numbers
		// 负数测试
		assert.Equal(t, -1, If(true, -1, 1))
		assert.Equal(t, -42, If(false, 42, -42))
	})

	t.Run("float type", func(t *testing.T) {
		// Basic float tests
		// 基础浮点数测试
		assert.Equal(t, 1.1, If(true, 1.1, 2.2))
		assert.Equal(t, 2.2, If(false, 1.1, 2.2))

		// Special float values
		// 特殊浮点数值测试
		assert.Equal(t, math.Inf(1), If(true, math.Inf(1), 0.0))
		assert.Equal(t, math.Inf(-1), If(false, 0.0, math.Inf(-1)))
		assert.True(t, math.IsNaN(If(true, math.NaN(), 1.0)))

		// Very small and large numbers
		// 极小和极大数值测试
		assert.Equal(t, math.SmallestNonzeroFloat64, If(true, math.SmallestNonzeroFloat64, 1.0))
		assert.Equal(t, math.MaxFloat64, If(false, 1.0, math.MaxFloat64))
	})

	t.Run("boolean type", func(t *testing.T) {
		assert.Equal(t, true, If(true, true, false))
		assert.Equal(t, false, If(false, true, false))
		// Test same value for both branches
		// 测试两个分支相同值的情况
		assert.Equal(t, true, If(true, true, true))
		assert.Equal(t, false, If(false, false, false))
	})

	t.Run("struct type", func(t *testing.T) {
		type testStruct struct {
			value string
			num   int
			valid bool
		}

		// Test with non-empty structs
		// 测试非空结构体
		s1 := testStruct{value: "struct1", num: 1, valid: true}
		s2 := testStruct{value: "struct2", num: 2, valid: false}
		assert.Equal(t, s1, If(true, s1, s2))
		assert.Equal(t, s2, If(false, s1, s2))

		// Test with empty struct
		// 测试空结构体
		emptyStruct := testStruct{}
		assert.Equal(t, emptyStruct, If(true, emptyStruct, s1))
		assert.Equal(t, emptyStruct, If(false, s1, emptyStruct))
	})

	t.Run("slice type", func(t *testing.T) {
		// Test with non-empty slices
		// 测试非空切片
		slice1 := []int{1, 2, 3}
		slice2 := []int{4, 5, 6}
		assert.Equal(t, slice1, If(true, slice1, slice2))
		assert.Equal(t, slice2, If(false, slice1, slice2))

		// Test with nil slice
		// 测试 nil 切片
		var nilSlice []int
		assert.Equal(t, nilSlice, If(true, nilSlice, slice1))
		assert.Equal(t, nilSlice, If(false, slice1, nilSlice))

		// Test with empty slice
		// 测试空切片
		emptySlice := make([]int, 0)
		assert.Equal(t, emptySlice, If(true, emptySlice, slice1))
		assert.NotEqual(t, nilSlice, emptySlice) // verify nil != empty
	})

	t.Run("map type", func(t *testing.T) {
		// Test with non-empty maps
		// 测试非空 map
		map1 := map[string]int{"a": 1, "b": 2}
		map2 := map[string]int{"c": 3, "d": 4}
		assert.Equal(t, map1, If(true, map1, map2))
		assert.Equal(t, map2, If(false, map1, map2))

		// Test with nil map
		// 测试 nil map
		var nilMap map[string]int
		assert.Equal(t, nilMap, If(true, nilMap, map1))
		assert.Equal(t, nilMap, If(false, map1, nilMap))

		// Test with empty map
		// 测试空 map
		emptyMap := make(map[string]int)
		assert.Equal(t, emptyMap, If(true, emptyMap, map1))
		assert.NotEqual(t, nilMap, emptyMap) // verify nil != empty
	})

	t.Run("channel type", func(t *testing.T) {
		// Test with buffered channels
		// 测试带缓冲的通道
		ch1 := make(chan int, 1)
		ch2 := make(chan int, 2)
		assert.Equal(t, ch1, If(true, ch1, ch2))
		assert.Equal(t, ch2, If(false, ch1, ch2))

		// Test with nil channel
		// 测试 nil 通道
		var nilChan chan int
		assert.Equal(t, nilChan, If(true, nilChan, ch1))
	})

	t.Run("time type", func(t *testing.T) {
		// Test with time values
		// 测试时间类型
		now := time.Now()
		future := now.Add(time.Hour)
		assert.Equal(t, now, If(true, now, future))
		assert.Equal(t, future, If(false, now, future))

		// Test with zero time
		// 测试零值时间
		zeroTime := time.Time{}
		assert.Equal(t, zeroTime, If(true, zeroTime, now))
	})

	t.Run("pointer type", func(t *testing.T) {
		// Test with valid pointers
		// 测试有效指针
		val1 := 42
		val2 := 24
		ptr1 := &val1
		ptr2 := &val2
		assert.Equal(t, ptr1, If(true, ptr1, ptr2))
		assert.Equal(t, ptr2, If(false, ptr1, ptr2))

		// Test with nil pointer
		// 测试 nil 指针
		var nilPtr *int
		assert.Equal(t, nilPtr, If(true, nilPtr, ptr1))
		assert.Equal(t, nilPtr, If(false, ptr1, nilPtr))
	})

	t.Run("interface type", func(t *testing.T) {
		// Test with interface values
		// 测试接口类型
		var i1 interface{} = "hello"
		var i2 interface{} = 42
		assert.Equal(t, i1, If(true, i1, i2))
		assert.Equal(t, i2, If(false, i1, i2))

		// Test with nil interface
		// 测试 nil 接口
		var nilInterface interface{}
		assert.Equal(t, nilInterface, If(true, nilInterface, i1))
	})
}
