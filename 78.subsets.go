/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 *
 * https://leetcode.cn/problems/subsets/description/
 *
 * algorithms
 * Medium (81.17%)
 * Total Accepted:    717.8K
 * Total Submissions: 884.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 * nums 中的所有元素 互不相同
 *
 *
 */
// Solution: Backtracking
func subsets(nums []int) [][]int {
	var result = make([]int, 0)
	var results = make([][]int, 0)
	results = subsets_dfs(nums, 0, result, results)
	return results
}

func subsets_dfs(nums []int, idx int, result []int, results [][]int) [][]int {
	var cpyResult = make([]int, len(result))
	copy(cpyResult, result)
	results = append(results, cpyResult)

	for ; idx < len(nums); idx++ {
		result = append(result, nums[idx])
		results = subsets_dfs(nums, idx+1, result, results)
		result = result[:len(result)-1]
	}
	return results
}

