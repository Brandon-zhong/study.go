package dp

import (
	"fmt"
	"math"
	"testing"
)

/**
给你k种面值的硬币，面值分别为c1, c2 ... ck，每种硬币的数量无限，再给一个总金额amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。算法的函数签名如下：
 coins 中是可选硬币面值，amount 是目标金额
int coinChange(int[] coins, int amount);
比如说k = 3，面值分别为 1，2，5，总金额amount = 11。那么最少需要 3 枚硬币凑出，即 11 = 5 + 5 + 1。
*/

/**
dp问题试解
dp三要素：重叠子问题、最优子结构、正确的「状态转移方程」
辅助思考状态转移方程:
明确（状态） -> 定义dp数组/函数的含义 -> 明确（选择）并择优 -> 明确 base case
*/

/**
状态就是原问题和子问题中变化的变量。这里硬币数量无限，变化的就是金额amount
dp函数的定义：函数dp(n)表示，当前目标金额是n，至少需要dp(n)个硬币凑出该金额
确定选择并择优：就是对每个状态，可以做出什么选择改变当前状态。在硬币问题上就是无论当前的目标金额是多少，选择就是从面额列表中选择一个硬币，然后目标金额相应减少。
*/

// 暴力递归 递归遍历所有情况
func coinChange1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt64
	for i := 0; i < len(coins); i++ {
		c := coinChange1(coins, amount-coins[i])
		if c == -1 {
			continue
		}
		res = int(math.Min(float64(c+1), float64(res)))
	}
	if res == math.MaxInt64 {
		return -1
	}
	return res
}

// 带备忘录的递归，重复的递归直接直接查备忘录
func coinChange2(coins []int, amount int) int {

	dict := make(map[int]int)
	var fn func(n int) int

	fn = func(n int) int {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		if i, ok := dict[n]; ok {
			return i
		}
		res := math.MaxInt64
		for i := 0; i < len(coins); i++ {
			c := fn(n - coins[i])
			if c == -1 {
				continue
			}
			res = int(math.Min(float64(c+1), float64(res)))

		}
		if res == math.MaxInt64 {
			dict[n] = -1
			return -1
		}
		dict[n] = res
		return res
	}

	return fn(amount)
}

// dp 数组的迭代解法,通过dp数组，记录下指定数值下的最小使用硬币数量，由于大的数额是由小数额组成的，通过小数额可以推导出大数额的组成数量
func coinChange3(coins []int, amount int) int {
	var dp []int = make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-coin])))
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func TestCoinChange1(t *testing.T) {
	fmt.Println(coinChange3([]int{1, 2, 5, 6}, 11))
}
func TestCoinChange2(t *testing.T) {
	fmt.Println(coinChange3([]int{1, 2, 5, 6}, 11))
}
func TestCoinChange3(t *testing.T) {
	fmt.Println(coinChange3([]int{1, 2, 5, 6}, 11))
}
