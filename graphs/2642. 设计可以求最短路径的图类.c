//go:build ignore

// 方法一：Dijkstra 求最短路径
typedef struct Edge {
    int to;
    int cost;
    struct Edge *next;
} Edge;
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

Edge *createEdge(int to, int cost) {
    Edge *obj = (Edge *)malloc(sizeof(Edge));
    obj->to = to;
    obj->cost = cost;
    obj->next = NULL;
    return obj;
}
void freeEdgeList(Edge *list) {
    while (list) {
        Edge *p = list;
        list = list->next;
        free(p);
    }
}
Node *createNode(long long x, int y) {
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
typedef struct {
    Edge **graph;
    int nodeSize;
} Graph;

Graph* graphCreate(int n, int** edges, int edgesSize, int* edgesColSize) {
    Graph *obj = (Graph *)malloc(sizeof(Graph));
    obj->nodeSize = n;
    obj->graph = (Edge **)malloc(sizeof(Edge *) * n);
    for (int i = 0; i < n; i++) {
        obj->graph[i] = NULL;
    }
    for (int i = 0; i < edgesSize; i++) {
        int x = edges[i][0];
        int y = edges[i][1];
        int cost = edges[i][2];
        Edge *e = (Edge *)malloc(sizeof(Edge));
        e->to = y;
        e->cost = cost;
        e->next = obj->graph[x];
        obj->graph[x] = e;
    }

    return obj;
}
void graphAddEdge(Graph* obj, int* edge, int edgeSize) {
    int x = edge[0];
    int y = edge[1];
    int cost = edge[2];
    Edge *e = (Edge *)malloc(sizeof(Edge));
    e->to = y;
    e->cost = cost;
    e->next = obj->graph[x];
    obj->graph[x] = e;
}
int graphShortestPath(Graph* obj, int node1, int node2) {
    int n = obj->nodeSize;
    int dist[n];
    PriorityQueue *pq = createPriorityQueue(n * n, greater);
    for (int i = 0; i < n; i++) {
        dist[i] = INT_MAX;
    }
    dist[node1] = 0;
    Node val;
    val.first = 0;
    val.second = node1;
    Push(pq, &val);
    while (!isEmpty(pq)) {
        Node *p = Pop(pq);
        int cost = p->first;
        int cur = p->second;
        if (cur == node2) {
            return cost;
        }
        for (Edge *pEntry = obj->graph[cur]; pEntry; pEntry = pEntry->next) {
            int next = pEntry->to;
            int ncost = pEntry->cost;
            if (dist[next] > cost + ncost) {
                dist[next] = cost + ncost;
                val.first = cost + ncost;
                val.second = next;
                Push(pq, &val);
            }
        }
    }
    return -1;
}
void graphFree(Graph* obj) {
    for (int i = 0; i < obj->nodeSize; i++) {
        freeEdgeList(obj->graph[i]);
    }
    free(obj->graph);
    free(obj);
}
