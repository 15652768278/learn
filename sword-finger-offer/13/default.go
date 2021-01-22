package main

var directGroup = [][]int{
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 0},
}

func movingCount(m int, n int, k int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	return dfs(m, n, 0, 0, k, dp)
}

func dfs(m, n, i, j, k int, dp [][]int) int {
	if i < 0 || j < 0 || i >= m || j >= n || dp[i][j] == 1 || sumPos(i)+sumPos(j) > k {
		return 0
	}
	dp[i][j] = 1
	ans := 1
	for _, direct := range directGroup {
		ans += dfs(m, n, i+direct[0], j+direct[1], k, dp)
	}
	return ans
}

func sumPos(n int) int {
	var ans int
	for n > 0 {
		ans += n % 10
		n = n / 10
	}
	return ans
}
