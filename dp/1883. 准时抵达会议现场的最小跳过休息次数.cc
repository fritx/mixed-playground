//go:build ignore

// 方法一：动态规划
// 注意 IEEE-754标准浮点数精度问题 可忽略误差处理
class Solution {
private:
    // 可忽略误差
    static constexpr double EPS = 1e-7;
    // 极大值
    static constexpr double INFTY = 1e20;

public:
    int minSkips(vector<int>& dist, int speed, int hoursBefore) {
        int n = dist.size();
        vector<vector<double>> f(n + 1, vector<double>(n + 1, INFTY));
        f[0][0] = 0.;
        for (int i = 1; i <= n; ++i) {
            for (int j = 0; j <= i; ++j) {
                if (j != i) {
                    f[i][j] = min(f[i][j], ceil(f[i - 1][j] + (double)dist[i - 1] / speed - EPS));
                }
                if (j != 0) {
                    f[i][j] = min(f[i][j], f[i - 1][j - 1] + (double)dist[i - 1] / speed);
                }
            }
        }
        for (int j = 0; j <= n; ++j) {
            if (f[n][j] < hoursBefore + EPS) {
                return j;
            }
        }
        return -1;
    }
};
