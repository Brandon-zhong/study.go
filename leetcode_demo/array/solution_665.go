package main

import (
	"sort"
)

/*
665. 非递减数列
给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。

我们是这样定义一个非递减数列的： 对于数组中所有的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

 

示例 1:

输入: nums = [4,2,3]
输出: true
解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
示例 2:

输入: nums = [4,2,1]
输出: false
解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
 

说明：

1 <= n <= 10 ^ 4
- 10 ^ 5 <= nums[i] <= 10 ^ 5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/non-decreasing-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//检测有一个降序对，即num[i] > nums[i + 1]，则要使这个数组递增，
//则要么使nums[i]降到nums[i + 1]的值，要么使nums[i + 1]升到nums[i]的值
//且修改后必须使nums处于非递减的状态，否则就无法在最多一次修改的情况下达成题意
func checkPossibility1(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			nums[i] = y
			if sort.IntsAreSorted(nums) {
				return true
			}
			nums[i] = x
			nums[i+1] = x
			return sort.IntsAreSorted(nums)
		}
	}
	return true
}

//上面的解法中，每次修改后需要遍历数组判断是否处于非递减状态，属于嵌套遍历
//要一次遍历完成检测的话，需要一次修改成功完成非递减状态，即使nums[i]降到nums[i + 1]的值后
//nums[i - 1] <= nums[i]才可以达成题意，否则尝试使nums[i + 1]升到nums[i]的值，如果前面的条件达成后
//继续向后遍历，是否还有num[i] > nums[i + 1]的情况，如果还有，则不满足题意，直接返false，否则遍历完成，返回true
func checkPossibility2(nums []int) bool {
	flag := false
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			if flag {
				return false
			}
			flag = true
			if i != 0 && nums[i-1] > y {
				nums[i+1] = x
			}
		}
	}
	return true
}

/*func main() {

	fmt.Println(checkPossibility2([]int{2, 3, 3, 2, 2}))

}*/
