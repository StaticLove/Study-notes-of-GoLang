package main

import (
	"fmt"
	"testing"
)

func Test_LastMonday(t *testing.T) {
	time := LastMonday()
	fmt.Printf("上一个周一日期为：%v\n", time)
}
