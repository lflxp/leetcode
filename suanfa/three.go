package main

import (
	"fmt"
	"strings"
)

/*
@title 3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
	 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
@分析
1.确定第一个
2.判断期间相邻的是否重复
3.计算下标长度
@结果
显示详情
执行用时 : 572 ms, 在Longest Substring Without Repeating Characters的Go提交中击败了4.58% 的用户
内存消耗 : 6.9 MB, 在Longest Substring Without Repeating Characters的Go提交中击败了19.59% 的用户
*/
func lengthOfLongestSubstring(s string) int {
	temp := strings.Split(s, "")
	total := 0  // 最长长度
	start := 0  // 当前位置
	circle := 0 // 循环次数
	count := map[string]int{}
Back:
	// fmt.Println("start ", start)
	for ; start < len(temp); start++ {
		if _, ok := count[temp[start]]; ok {
			// fmt.Println("exist ", temp[start], count)
			if len(count) > total {
				total = len(count)
			}
			circle++
			start = circle
			count = map[string]int{}
			goto Back
		} else {
			count[temp[start]] = 1
			// fmt.Println("not repet ", count, temp[start], start)
		}
	}
	if len(count) > total {
		// fmt.Println("last count total", count, len(count), total)
		total = len(count)
	}
	return total
}

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
