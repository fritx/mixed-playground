//go:build ignore

class Solution {
public:
    string capitalizeTitle(string title) {
        int n = title.size();
        int l = 0, r = 0;   // 单词左右边界（左闭右开）
        while (r < n) {
            while (r < n && title[r] != ' ') {
                ++r;
            }
            // 对于每个单词按要求处理
            if (r - l > 2) {
                title[l++] = toupper(title[l]);
            }
            while (l < r) {
                title[l++] = tolower(title[l]);
            }
            l = ++r;
        }
        return title;
    }
};
