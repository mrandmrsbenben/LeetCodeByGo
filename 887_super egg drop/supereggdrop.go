package main

import "fmt"

func main() {
	fmt.Println("Output: ", superEggDrop(3, 14))
}

func superEggDrop(K int, N int) int {
	if N == 1 {
		return 1
	}

	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
	}
	for i := 1; i <= K; i++ {
		dp[1][i] = 1
	}

	res := -1
	for i := 2; i <= N; i++ {
		for j := 1; j <= K; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j] + 1
		}
		if dp[i][K] >= N {
			res = i
			break
		}
	}
	return res
}

// public int superEggDrop(int K, int N) {
//     if (N == 1) {
//         return 1;
//     }
//     int[][] f = new int[N + 1][K + 1];
//     for (int i = 1; i <= K; ++i) {
//         f[1][i] = 1;
//     }
//     int ans = -1;
//     for (int i = 2; i <= N; ++i) {
//         for (int j = 1; j <= K; ++j) {
//             f[i][j] = 1 + f[i - 1][j - 1] + f[i - 1][j];
//         }
//         if (f[i][K] >= N) {
//             ans = i;
//             break;
//         }
//     }
//     return ans;
// }
