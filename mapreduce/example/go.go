package example

import (
	"log"
	"runtime/debug"
)

// Go recover 野生goroutine
// 可以替换go关键字使用 有参数函数使用闭包
func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				stack := string(debug.Stack())
				log.Println(stack)
			}
		}()
		f()
	}()
}
