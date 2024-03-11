/*
 * @lc app=leetcode id=1524 lang=golang
 *
 * [1524] Number of Sub-arrays With Odd Sum
 *
 * https://leetcode.cn/problems/number-of-sub-arrays-with-odd-sum/description/
 *
 * algorithms
 * Medium (46.92%)
 * Total Accepted:    10.8K
 * Total Submissions: 22.8K
 * Testcase Example:  '[1,3,5]'
 *
 * 给你一个整数数组 arr 。请你返回和为 奇数 的子数组数目。
 *
 * 由于答案可能会很大，请你将结果对 10^9 + 7 取余后返回。
 *
 *
 *
 * 示例 1：
 *
 * 输入：arr = [1,3,5]
 * 输出：4
 * 解释：所有的子数组为 [[1],[1,3],[1,3,5],[3],[3,5],[5]] 。
 * 所有子数组的和为 [1,4,9,3,8,5].
 * 奇数和包括 [1,9,3,5] ，所以答案为 4 。
 *
 *
 * 示例 2 ：
 *
 * 输入：arr = [2,4,6]
 * 输出：0
 * 解释：所有子数组为 [[2],[2,4],[2,4,6],[4],[4,6],[6]] 。
 * 所有子数组和为 [2,6,12,4,10,6] 。
 * 所有子数组和都是偶数，所以答案为 0 。
 *
 *
 * 示例 3：
 *
 * 输入：arr = [1,2,3,4,5,6,7]
 * 输出：16
 *
 *
 * 示例 4：
 *
 * 输入：arr = [100,100,99,99]
 * 输出：4
 *
 *
 * 示例 5：
 *
 * 输入：arr = [7]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 10^5
 * 1 <= arr[i] <= 100
 *
 *
 */

const MASK = 1e+09 + 7

// dp(i, j): Sum of subarray [i..j] is odd
//
// dp(i, i) = true,  arr[i] is odd
//            false, arr[i] is not odd
// dp(i, j) = true,  (dp(i, j-1) && !dp(j, j)) or (!dp(i, j-1) && dp(j, j))
//            false, (dp(i, j-1) && dp(j, j)) or (!dp(i, j-1) && !dp(j, j))
func numOfSubarrays(arr []int) int {
	count := 0
	n := len(arr)

	var dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)

		// Initialize dp(i, i)
		if arr[i]%2 != 0 {
			dp[i][i] = true
			count++
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (dp[i][j-1] && !dp[j][j]) ||
				(!dp[i][j-1] && dp[j][j]) {
				dp[i][j] = true
				count++
			} else {
				dp[i][j] = false
			}
		}
	}
	return count % MASK
}
