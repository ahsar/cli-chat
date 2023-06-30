package record

import (
	"fmt"
	"testing"
)

func TestContainer(t *testing.T) {
	buffer := NewCircularBuffer(5)

	// 添加元素到缓冲区
	buffer.Append("1")
	buffer.Append("2")
	buffer.Append("3")
	buffer.Append("4")
	buffer.Append("5")

	buffer.Append("6")
	buffer.Append("7")
	buffer.Append("8")
	buffer.Append("9")
	buffer.Append("10")
	buffer.Append("11")
	buffer.Append("12")
	buffer.Append("13")
	buffer.Append("14")

	// 获取最早添加的元素
	var earliest string
	if earliest = buffer.GetEarliest(); earliest != "10" {
		t.Fail()
	}
	fmt.Println("Earliest element:", earliest)

	// 从最早元素到最后元素遍历输出
	elements := buffer.Iterate()
	fmt.Println("Elements:", elements)
}
