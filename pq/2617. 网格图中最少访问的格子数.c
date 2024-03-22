//go:build ignore

typedef struct {
    int first;
    int second;
} Node;
typedef bool (*cmp)(const void *, const void *);

typedef struct {
    Node *arr;
    int capacity;
    int queueSize;
    cmp compare;
} PriorityQueue;

Node *createNode(int x, int y) {
    Node *obj = (Node *)malloc(sizeof(Node));
    obj->first = x;
    obj->second = y;
    return obj;
}
PriorityQueue *createPriorityQueue(int size, cmp compare) {
    PriorityQueue *obj = (PriorityQueue *)malloc(sizeof(PriorityQueue));
    obj->arr = (Node *)malloc(sizeof(Node) * size);
    obj->queueSize = 0;
    obj->capacity = size;
    obj->compare = compare;
    return obj;
}
static void swap(Node *arr, int i, int j) {
    Node tmp;
    memcpy(&tmp, &arr[i], sizeof(Node));
    memcpy(&arr[i], &arr[j], sizeof(Node));
    memcpy(&arr[j], &tmp, sizeof(Node));
}
static void down(Node *arr, int size, int i, cmp compare) {
    for (int k = 2 * i + 1; k < size; k = 2 * k + 1) {
        // 父节点 (k - 1) / 2，左子节点 k，右子节点 k + 1
        if (k + 1 < size && compare(&arr[k], &arr[k + 1])) {
            k++;
        }
        if (compare(&arr[k], &arr[(k - 1) / 2])) {
            break;
        }
        swap(arr, k, (k - 1) / 2);
    }
}
void Heapify(PriorityQueue *obj) {
    for (int i = obj->queueSize / 2 - 1; i >= 0; i--) {
        down(obj->arr, obj->queueSize, i, obj->compare);
    }
}
void Push(PriorityQueue *obj, Node *node) {
    memcpy(&obj->arr[obj->queueSize], node, sizeof(Node));
    for (int i = obj->queueSize; i > 0 && obj->compare(&obj->arr[(i - 1) / 2], &obj->arr[i]); i = (i - 1) / 2) {
        swap(obj->arr, i, (i - 1) / 2);
    }
    obj->queueSize++;
}
Node* Pop(PriorityQueue *obj) {
    swap(obj->arr, 0, obj->queueSize - 1);
    down(obj->arr, obj->queueSize - 1, 0, obj->compare);
    Node *ret =  &obj->arr[obj->queueSize - 1];
    obj->queueSize--;
    return ret;
}
bool isEmpty(PriorityQueue *obj) {
    return obj->queueSize == 0;
}
Node* Top(PriorityQueue *obj) {
    if (obj->queueSize == 0) {
        return NULL;
    } else {
        return &obj->arr[0];
    }
}
void FreePriorityQueue(PriorityQueue *obj) {
    free(obj->arr);
    free(obj);
}
bool greater(const void *a, const void *b) {
    return ((Node *)a)->first > ((Node *)b)->first;
}
int update(int x, int y) {
    if (x == -1 || y < x) {
        return y;
    } else {
        return x;
    }
}
int minimumVisitedCells(int** grid, int gridSize, int* gridColSize){
    int m = gridSize, n = gridColSize[0];
    int dist[m][n];
    for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
            dist[i][j] = -1;
        }
    }
    dist[0][0] = 1;
    PriorityQueue *row[m], *col[n];
    for (int i = 0; i < m; i++) {
        row[i] = createPriorityQueue(n, greater);
    }
    for (int i = 0; i < n; i++) {
        col[i] = createPriorityQueue(m, greater);
    }
    for (int i = 0; i < m; ++i) {
        for (int j = 0; j < n; ++j) {
            while (!isEmpty(row[i]) && Top(row[i])->second + grid[i][Top(row[i])->second] < j) {
                Pop(row[i]);
            }
            if (!isEmpty(row[i])) {
                dist[i][j] = update(dist[i][j], dist[i][Top(row[i])->second] + 1);
            }
            while (!isEmpty(col[j]) && Top(col[j])->second + grid[Top(col[j])->second][j] < i) {
                Pop(col[j]);
            }
            if (!isEmpty(col[j])) {
                dist[i][j] = update(dist[i][j], dist[Top(col[j])->second][j] + 1);
            }
            if (dist[i][j] != -1) {
                Node node;
                node.first = dist[i][j];
                node.second = j;
                Push(row[i], &node);
                node.second = i;
                Push(col[j], &node);
            }
        }
    }
    return dist[m - 1][n - 1];
}
