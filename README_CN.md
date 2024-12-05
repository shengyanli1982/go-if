[English](./README.md) | 中文

<div align="center">
	<h1>go-if</h1>
    <p>一个轻量优雅的 Go 语言条件表达式实现 - 你一直期待的三元运算符</p>
    <img src="assets/logo.png" alt="logo" width="350px">
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/shengyanli1982/go-if)](https://goreportcard.com/report/github.com/shengyanli1982/go-if)
[![Build Status](https://github.com/shengyanli1982/go-if/actions/workflows/test.yaml/badge.svg)](github.com/shengyanli1982/go-if/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/shengyanli1982/go-if.svg)](https://pkg.go.dev/github.com/shengyanli1982/go-if)

# 简介

在 Go 中想念三元运算符？来试试 `go-if` 吧！这个轻量级库为 Go 带来了优雅的条件表达式，让你的代码更简洁、更富有表现力。它就像三元运算符，但以一种更符合 Go 语言风格的方式实现。

`go-if` 提供：

1. 支持任意类型的泛型实现
2. 清晰简单、易于理解的语法
3. 零外部依赖
4. 全面的测试覆盖
5. 类型安全的操作
6. 完美支持单行表达式，且不影响可读性

# 为什么选择 go-if？

-   **简单的 API**：一个函数完美实现一个功能
-   **类型安全**：完整的泛型实现，支持任意类型
-   **零依赖**：不需要任何外部包
-   **生产就绪**：经过各种类型和边界情况的充分测试
-   **熟悉的模式**：用法类似于你熟悉的三元运算符
-   **Go 语言风格**：设计符合 Go 代码的自然风格

# 安装

使用 `go get` 命令安装 `go-if`：

```bash
go get github.com/shengyanli1982/go-if
```

# 快速开始

`go-if` 的使用非常简单：

```go
package main

import (
    "fmt"
    goif "github.com/shengyanli1982/go-if"
)

func main() {
    // 简单的字符串条件
    result := goif.If(true, "yes", "no")
    fmt.Println(result) // 输出: yes

    // 数字操作
    count := goif.If(len(items) > 5, 10, 5)
    fmt.Println(count)

    // 复杂类型
    user := goif.If(isAdmin,
        User{Role: "admin", Permissions: []string{"read", "write"}},
        User{Role: "guest", Permissions: []string{"read"}},
    )
}
```

# 特性和用例

## 1. 基本类型操作

完美适用于简单的条件赋值：

```go
// 字符串操作
greeting := goif.If(isEnglish, "Hello", "你好")

// 数值操作
discount := goif.If(isPremiumUser, 0.2, 0.1)

// 布尔标志
isEnabled := goif.If(config.Debug, true, false)
```

## 2. 复杂类型操作

优雅处理任意类型：

```go
// 结构体选择
config := goif.If(isProduction,
    Config{
        Host:    "prod.example.com",
        Timeout: 30,
    },
    Config{
        Host:    "dev.example.com",
        Timeout: 5,
    },
)

// 切片操作
permissions := goif.If(isAdmin,
    []string{"read", "write", "delete"},
    []string{"read"},
)

// 映射选择
cache := goif.If(isDistributed,
    map[string]int{"capacity": 1000, "ttl": 3600},
    map[string]int{"capacity": 100, "ttl": 300},
)
```

## 3. 指针和接口处理

安全地处理指针和接口：

```go
// 指针选择
var fallbackCache *Cache = nil
activeCache := goif.If(isCacheEnabled,
    mainCache,
    fallbackCache,
)

// 接口处理
var handler interface{} = goif.If(isAsync,
    AsyncHandler{},
    SyncHandler{},
)
```

## 4. 时间和持续时间操作

完美适用于时间相关的条件：

```go
// 超时时间
timeout := goif.If(isHighPriority,
    5 * time.Second,
    30 * time.Second,
)

// 时间选择
deadline := goif.If(isUrgent,
    time.Now().Add(1 * time.Hour),
    time.Now().Add(24 * time.Hour),
)
```

# 性能基准测试

我们对 `go-if` 和传统 if-else 语句在不同数据类型下进行了全面的性能对比测试：

```bash
$ go test -bench=. -benchmem
goos: windows
goarch: amd64
pkg: github.com/shengyanli1982/go-if
cpu: 12th Gen Intel(R) Core(TM) i5-12400F
BenchmarkString/Generic_If-12                   1000000000               0.1293 ns/op          0 B/op          0 allocs/op
BenchmarkString/Traditional_IfElse-12           1000000000               0.1297 ns/op          0 B/op          0 allocs/op
BenchmarkInteger/Generic_If-12                  1000000000               0.1299 ns/op          0 B/op          0 allocs/op
BenchmarkInteger/Traditional_IfElse-12          1000000000               0.1310 ns/op          0 B/op          0 allocs/op
BenchmarkStruct/Generic_If-12                   1000000000               0.1285 ns/op          0 B/op          0 allocs/op
BenchmarkStruct/Traditional_IfElse-12           1000000000               0.1306 ns/op          0 B/op          0 allocs/op
BenchmarkSlice/Generic_If-12                    1000000000               0.1284 ns/op          0 B/op          0 allocs/op
BenchmarkSlice/Traditional_IfElse-12            1000000000               0.1281 ns/op          0 B/op          0 allocs/op
```

## 运行基准测试

要自行运行基准测试，请执行：

```bash
go test -bench=. -benchmem
```

针对特定类型的基准测试：

```bash
# 仅测试字符串操作
go test -bench=BenchmarkString -benchmem

# 仅测试整数操作
go test -bench=BenchmarkInteger -benchmem

# 仅测试结构体操作
go test -bench=BenchmarkStruct -benchmem

# 仅测试切片操作
go test -bench=BenchmarkSlice -benchmem
```

## 基准测试用例

基准测试涵盖了四个主要类别：

1. **字符串操作**：比较字符串条件赋值的性能
2. **整数操作**：测试数值选择的性能
3. **结构体操作**：评估复杂结构体类型处理的性能
4. **切片操作**：测量切片赋值的性能

每个基准测试都比较了：

-   `Generic_If`：使用我们的泛型 `If` 函数
-   `Traditional_IfElse`：使用传统的 if-else 语句

这些测试结果可以帮助你在项目中更好地决定何时使用 `go-if`。

# 最佳实践

## 1. 保持简单

`go-if` 的优美之处在于它的简单性。请用它处理清晰、直接的条件：

```go
// 好 - 清晰简单
message := goif.If(count > 0, "找到项目", "无项目")

// 避免 - 过于复杂
// 对于复杂逻辑，应使用常规的 if 语句
result := goif.If(
    complexCondition() && anotherCondition(),
    someComplexFunction(),
    anotherComplexFunction(),
)
```

## 2. 类型一致性

确保两个返回值的类型相同：

```go
// 好 - 两个值都是字符串
path := goif.If(isWindows, "C:\\path", "/path")

// 差 - 混合类型将无法编译
// value := goif.If(condition, "string", 123)
```

# 局限性

让我们明确 `go-if` 不适用的场景：

-   不能替代所有的 if 语句（复杂逻辑应使用常规 if 语句）
-   不适合有副作用的操作（需要执行多个语句时应使用常规 if 语句）
-   不适合嵌套条件（考虑重构或使用常规 if 语句）

# 贡献

欢迎为 `go-if` 做出贡献！请随时提交 Pull Request。

# 许可证

`go-if` 基于 MIT 许可证发布。详情请查看 LICENSE 文件。
