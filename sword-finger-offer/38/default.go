package main

import (
	"sort"
	"strings"
)

func permutation(s string) []string {
	if len(s) == 0 {
		return nil
	}
	ans := make([]string, 0)
	str := strings.Split(s, "")
	sort.Strings(str)
	dfs(str, 0, &ans)
	return ans
}

func dfs(str []string, i int, ans *[]string) {
	n := len(str)
	for i == n-1 {
		*ans = append(*ans, strings.Join(str, ""))
		return
	}
	for k := i; k < n; k++ {
		if k != i && str[k] == str[i] {
			continue
		}
		str[i], str[k] = str[k], str[i]
		dfs(str, i+1, ans)
	}
	for k := n - 1; k > i; k-- {
		str[i], str[k] = str[k], str[i]
	}
}
