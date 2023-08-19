package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Print("请输入溶液种类数量n：")
	var n int
	fmt.Scan(&n)

	fmt.Print("请输入化学反应增加单位x：")
	var x int
	fmt.Scan(&x)

	fmt.Print("请输入需要达到的体积c：")
	var c int
	fmt.Scan(&c)

	v := make([]int, n)
	w := make([]int, n)

	fmt.Println("现在请依次输入每种溶液的体积和物质含量：")
	for i := 0; i < n; i++ {
		fmt.Printf("请输入第%d种溶液的体积v[%d]：", i+1, i)
		fmt.Scan(&v[i])

		fmt.Printf("请输入第%d种溶液的物质含量w[%d]：", i+1, i)
		fmt.Scan(&w[i])
	}

	// dp数组，用于存储每个体积的最大物质含量
	dp := make([]int, c+1)

	// 遍历每种溶液
	for i := 0; i < n; i++ {
		// 更新每个体积的物质含量
		for j := v[i]; j <= c; j++ {
			dp[j] = max(dp[j], dp[j-v[i]]+w[i])
		}
	}

	// 计算同体积合并后的物质含量
	for i := 1; i <= c; i++ {
		dp[i] = max(dp[i], dp[i-1]+x)
	}

	fmt.Println("物质含量最多是：", dp[c])
}
