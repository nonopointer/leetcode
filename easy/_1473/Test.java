package easy._1473;

public class Test {
    public static void main(String[] args) {
        
    }

    /**
     * 输入：houses = [0,0,0,0,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n
     * = 2, target = 3 输出：9 解释：房子涂色方案为 [1,2,2,1,1] 此方案包含 target = 3 个街区， 分别是
     * [{1},{2,2}, {1,1}]。 涂色的总花费为 (1 + 1 + 1 + 1 + 5) = 9。
     * 
     */
    static final int INF = 0x3f3f3f3f; // INF + INF 不会爆int
    static int[][][] f = new int[105][25][105];

    public int minCost(int[] houses, int[][] cost, int m, int n, int target) {
        // 无效的状态
        for (int i = 0; i <= m; i++) {
            for (int j = 0; j <= n; j++) {
                f[i][j][0] = INF;
            }
        }

        for (int i = 1; i <= m; i++) {
            int color = houses[i - 1];
            for (int j = 1; j <= n; j++) {
                for (int k = 1; k <= target; k++) {
                    if (k > i) {
                        f[i][j][k] = INF;
                        continue;
                    }
                    // 第i个房间已经上色
                    if (color != 0) {
                        if (j == color) {
                            int cur = INF;
                            // 与上一个房间颜色不同
                            for (int p = 1; p <= n; p++) {
                                if (p != j) { // 颜色不同
                                    cur = Math.min(cur, f[i - 1][p][k - 1]);
                                }
                            }
                            // 与上一个房间颜色相同
                            f[i][j][k] = Math.min(cur, f[i - 1][j][k]);
                        } else
                            f[i][j][k] = INF; // 其他为无效状态
                    } else { // 第i个房间未上色
                        int u = cost[i - 1][j - 1];
                        // 考虑第i个房间是否形成新的分区
                        // 与上一个房间颜色不同，形成新分区
                        int cur = INF;
                        for (int p = 1; p <= n; p++) {
                            if (p != j)
                                cur = Math.min(cur, f[i - 1][p][k - 1] + u);
                        }
                        // 与上一个房间颜色相同，不形成新的分区
                        f[i][j][k] = Math.min(cur, f[i - 1][j][k] + u);
                    }
                }
            }
        }

        int ans = INF;
        for (int i = 1; i <= n; i++) {
            ans = Math.min(ans, f[m][i][target]);
        }
        return ans == INF ? -1 : ans;
    }
}