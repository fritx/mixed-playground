
//go:build ignore

// 方法二：Floyd 求最短路径
typedef struct {
    int **dist;
    int n;
} Graph;

Graph* graphCreate(int n, int** edges, int edgesSize, int* edgesColSize) {
    Graph *obj = (Graph *)malloc(sizeof(Graph));
    obj->dist = (int **)malloc(sizeof(int *) * n);
    obj->n = n;
    for (int i = 0; i < n; i++) {
        obj->dist[i] = (int *)malloc(sizeof(int) * n);
        for (int j = 0; j < n; j++) {
            obj->dist[i][j] = INT_MAX;
        }
         obj->dist[i][i] = 0;
    }
    for (int i = 0; i < edgesSize; i++) {
        int x = edges[i][0], y = edges[i][1], cost = edges[i][2];
        obj->dist[x][y] = cost;
    }
    for (int k = 0; k < n; k++) {
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (obj->dist[i][k] != INT_MAX && obj->dist[k][j] != INT_MAX) {
                    obj->dist[i][j] = fmin(obj->dist[i][j], obj->dist[i][k] + obj->dist[k][j]);
                }
            }
        }
    }
    return obj;
}
void graphAddEdge(Graph* obj, int* edge, int edgeSize) {
    int x = edge[0], y = edge[1], cost = edge[2];
    if (cost >= obj->dist[x][y]) {
        return;
    }
    for (int i = 0; i < obj->n; i++) {
        for (int j = 0; j < obj->n; j++) {
            if (obj->dist[i][x] != INT_MAX && obj->dist[y][j] != INT_MAX) {
                obj->dist[i][j] = fmin(obj->dist[i][j], obj->dist[i][x] + cost + obj->dist[y][j]);
            }
        }
    }
}
int graphShortestPath(Graph* obj, int node1, int node2) {
    int res = obj->dist[node1][node2];
    return res == INT_MAX ? -1 : res;
}
void graphFree(Graph* obj) {
    for (int i = 0; i < obj->n; i++) {
        free(obj->dist[i]);
    }
    free(obj->dist);
    free(obj);
}
