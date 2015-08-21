package logger_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	"time"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("logger", func() {

		fmt.Println("Testing")
	})
})

func funcs() {
	for i := 1; i < 10; i = i + 1 {
		go func() {
			time.Sleep(time.Second)
			fmt.Println("funcs go " + time.Now().String())
		}()
	}
}

func funcs1() {

	for i := 1; i < 10; i = i + 1 {
		time.Sleep(1 * time.Second)
		fmt.Println("funcs1 go " + time.Now().String())
	}
}
