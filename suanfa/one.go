package main

import "fmt"

/*
@title 两数之和
@result
成功
显示详情
执行用时 : 60 ms, 在Two Sum的Go提交中击败了33.60% 的用户
内存消耗 : 3.1 MB, 在Two Sum的Go提交中击败了41.54% 的用户
进行下一个挑战：
三数之和
四数之和
两数之和 II - 输入有序数组
Two Sum III - Data structure design
和为K的子数组
两数之和 IV - 输入 BST
炫耀一下:

*/
func twoSum(nums []int, target int) []int {
	rs := []int{}
	for i, v := range nums {
		for i1, v1 := range nums {
			if i1 > i {
				// fmt.Println(i, i1, v, v1, v+v1)
				if v+v1 == target {
					// fmt.Println("got it ", i, i1, v, v1)
					rs = append(rs, i, i1)
					goto STOP
				}
			}
		}
	}
STOP:
	return rs
}

func main() {
	tmp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	target := 12
	fmt.Println(twoSum(tmp, target))
}
