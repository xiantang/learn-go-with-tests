package array

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		// 如果你使用[4]int 作为参数传入函数
		// 不能通过编译 因为他们是不同的类型
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		// %v 用来展示数组
		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
