package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type testFn func(int)

func main() {
	f := withLogging(test)
	f(10)

	f1 := wrapLogger(test, log.New(os.Stdout, "test: ", 1))
	f1(3)
}

func withLogging(fn testFn) testFn {
	return func(n int) {
		fmt.Println("before")
		now := time.Now()
		fn(n)
		fmt.Printf("after (%v)\n", time.Since(now))
	}
}
func test(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func wrapLogger(do testFn, logger *log.Logger) testFn {
	return func(n int) {
		t := time.Now()
		do(n)
		logger.Printf("took=%v, n=%v", time.Since(t), n)
	}
}

// func wrapLogger(do testFn, logger *log.Logger) testFn {
// 	return func(n int) {
// 		fn := func(n int) {
// 			defer func(t time.Time) {
// 				logger.Printf(
// 					"took=%v, n=%v",
// 					time.Since(t),
// 					n)
// 			}(time.Now())

// 			do(n)
// 		}
// 		fn(n)
// 	}
// }
