package main

import (
	"fmt"
	"strings"
)

/*
58. 最后一个单词的长度
给你一个字符串 s，由若干单词组成，单词之间用空格隔开。返回字符串中最后一个单词的长度。如果不存在最后一个单词，请返回 0 。

单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。

 

示例 1：

输入：s = "Hello World"
输出：5
示例 2：

输入：s = " "
输出：0
 

提示：

1 <= s.length <= 104
s 仅有英文字母和空格 ' ' 组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/length-of-last-word
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//双指针法，trim掉前后的空格，第一指针停在最后的一个空格处，第二指针尝试下一个空格
//找到下一个空格，则更新第一指针，直到遍历结束，根据第一指针和第二指针的距离得出最后单词长度
func lengthOfLastWord1(s string) int {
	s = strings.Trim(s, " ")
	first, second := 0, 0
	for ; second < len(s); second++ {
		if s[second] == ' ' {
			first = second
		}
	}
	if first != len(s) && s[first] == ' ' {
		first++
	}
	return second - first
}

//指针法，不使用trim，
func lengthOfLastWord2(s string) int {
	return 0
}

func main() {
	fmt.Println(lengthOfLastWord1("a"))

}
