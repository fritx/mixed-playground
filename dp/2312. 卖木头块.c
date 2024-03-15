//go:build ignore

typedef struct {
    long long key;
    long long val;
    UT_hash_handle hh;
} HashItem;

HashItem *hashFindItem(HashItem **obj, long long key) {
    HashItem *pEntry = NULL;
    HASH_FIND(hh, *obj, &key, sizeof(long long), pEntry);
    return pEntry;
}

bool hashAddItem(HashItem **obj, long long key, long long val) {
    if (hashFindItem(obj, key)) {
        return false;
    }
    HashItem *pEntry = (HashItem *)malloc(sizeof(HashItem));
    pEntry->key = key;
    pEntry->val = val;
    HASH_ADD(hh, *obj, key, sizeof(key), pEntry);
    return true;
}

bool hashSetItem(HashItem **obj, long long key, long long val) {
    HashItem *pEntry = hashFindItem(obj, key);
    if (!pEntry) {
        hashAddItem(obj, key, val);
    } else {
        pEntry->val = val;
    }
    return true;
}

int hashGetItem(HashItem **obj, long long key, long long defaultVal) {
    HashItem *pEntry = hashFindItem(obj, key);
    if (!pEntry) {
        return defaultVal;
    }
    return pEntry->val;
}

void hashFree(HashItem **obj) {
    HashItem *curr = NULL, *tmp = NULL;
    HASH_ITER(hh, *obj, curr, tmp) {
        HASH_DEL(*obj, curr);
        free(curr);
    }
}

long long pairHash(int x, int y) {
    return ((long long) x << 16) ^ y;
}

long long dfs(int x, int y, long long **memo, HashItem **val) {
    if (memo[x][y] != -1) {
        return memo[x][y];
    }

    long long ret = hashGetItem(val, pairHash(x, y), 0LL);
    if (x > 1) {
        for (int i = 1; i < x; ++i) {
            ret = fmax(ret, dfs(i, y, memo, val) + dfs(x - i, y, memo, val));
        }
    }
    if (y > 1) {
        for (int j = 1; j < y; ++j) {
            ret = fmax(ret, dfs(x, j, memo, val) + dfs(x, y - j, memo, val));
        }
    }
    return memo[x][y] = ret;
};

long long sellingWood(int m, int n, int** prices, int pricesSize, int* pricesColSize) {
    HashItem *value = NULL;
    long long *memo[m + 1];
    for (int i = 0; i <= m; i++) {
        memo[i] = (long long *)malloc(sizeof(long long) * (n + 1));
        for (int j = 0; j <= n; j++) {
            memo[i][j] = -1;
        }
    }
    for (int i = 0; i < pricesSize; ++i) {
        hashAddItem(&value, pairHash(prices[i][0], prices[i][1]), prices[i][2]);
    }
    long long ret = dfs(m, n, memo, &value);
    for (int i = 0; i <= m; i++) {
        free(memo[i]);
    }
    hashFree(&value);
    return ret;
}
