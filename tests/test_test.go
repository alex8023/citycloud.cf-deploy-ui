package tests

import (
	"fmt"
	"strings"
	"testing"
)

func TestArr(testing *testing.T) {
	arr := []string{"192.168.138.10", "192.168.138.10", "192.168.138.10", "192.168.138.10"}

	s := strings.Join(arr, ",")
	fmt.Println(s)
}
