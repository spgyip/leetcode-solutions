/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.cn/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (37.97%)
 * Total Accepted:    1.6M
 * Total Submissions: 4.2M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 * 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 仅由数字和英文字母组成
 *
 *
 */

// Dynamic programing
//
//	 dp(i, j) = true, dp(i+1, j-1)==true and s[i]==s[j]
//		      = false, dp(i+1, j-1)==false or s[i]!=s[j]
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	start, maxL := 0, 1
	for wnd := 2; wnd <= n; wnd++ {
		for i := 0; i <= n-wnd; i++ {
			j := i + wnd - 1
			if s[i] != s[j] {
				dp[i][j] = 0
				continue
			}
			if wnd == 2 || dp[i+1][j-1] == 1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}

			if dp[i][j] == 1 && maxL < j-i+1 {
				maxL = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+maxL]
}

// Strategy 2: Expand from center
func longestPalindrome2(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		n0 := expandPalindrome(s, i, i)
		n1 := expandPalindrome(s, i, i+1)
		n := max(n0, n1)

		if n > end-start+1 {
			start = i - (n-1)/2
			end = start + n - 1
		}
	}
	return s[start : end+1]
}

func expandPalindrome(s string, left int, right int) int {
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		left, right = left-1, right+1
	}
	return right - left - 1
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

