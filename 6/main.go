package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    if n == 0 {
        return 0
    }

    dp := make([]int, n+1)
    dp[0] = 0
    dp[1] = 0

    for i := 2; i <= n; i++ {
        dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
    }

    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    cost := []int{10, 15, 20}
    fmt.Println("Minimum cost to reach the top:", minCostClimbingStairs(cost)) // Output: 15
}

// This solution has a time complexity of O(n) and a space complexity of O(n), where n is the length of the input array cost.