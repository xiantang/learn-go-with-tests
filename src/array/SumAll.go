package array

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	//make 可以在创建切片的时候指定我们需要的长度和容量。
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// 可以使用 = 对切片元素进行赋值
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums

}
