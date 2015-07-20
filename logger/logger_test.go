package logger

import (
	"fmt"
	"testing"
	"time"
)

func TestLogger(test *testing.T) {
	fmt.Println("Testing")

}
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
