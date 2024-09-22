
# Cheatsheet among languages

## Data Structures

Static Arrays:

| Lang  | Array | fill | copy | free |
| -- | -- | -- | -- | -- |
| JS    | a = new Array(3) | a.fill(0) | - | - |
| Python| a = list() | a = [0] * n | - | - |
| Go    | var a [3]int = [3]int{} | - | copy(newArray, fixedArray) | - |
| Java  | int[] a = new int[3] - | Arrays.fill(a, 0) | int[] newArray = new int[fixedArray.length * 2];<br>System.arraycopy(fixedArray, 0, newArray, 0, fixedArray.length); | - |
| C++   | int a[3]; | std::fill(std::begin(arr), std::end(arr), 0) / std::fill(arr, arr + 10, 0) | std::copy(std::begin(fixedArray), std::end(fixedArray), std::begin(newArray)) | delete[] a |
| C   | ↑ | memset(a, 0, sizeof(a)) | Node tmp; memcpy(&tmp, &arr[i], sizeof(Node)); memcpy(&arr[i], &arr[j], sizeof(Node)); | free a |

Dynamic Arrays / Slices / Vectors:

| Lang  | Slice | fill |
| -- | -- | -- |
| JS    | 同Array | ← |
| Python| 同Array | ← |
| Go    | var a []int = []int{} | - |
| Rust  | let a: Vec\<i32\> = vec![0; n] / Vec::new() | let vec: Vec\<i32\> = std::iter::repeat(0).take(10).collect() |
| Java  | List\<Integer\> a = new ArrayList\<Integer\>() / List\<String\> a = Arrays.asList("a", "b", "c") | List\<Integer\> list = IntStream.range(0, 10).mapToObj(i -> 0).collect(Collectors.toList()) / List\<Integer\> list = Collections.nCopies(10, 0) |
| C++   | std::vector\<int\> a = {1, 2, 3} | std::vector\<int\> vec(10, 0); / std::fill(vec.begin(), vec.end(), 0) |
| C     | char *list[] = {"a", "b", "c"} / int \*a = NULL; int size = 0; capacity = n; a = (int \*)malloc(capacity \* sizeof(int)); if (dynamicArray == NULL) return 1; | 同Array |

Slices Realloc:

| Lang  | Realloc |
| -- | -- |
| Go  | newSlice := make([]int, 0, 2*cap(s))<br>newSlice = append(newSlice, s...) |
| Java  | list.ensureCapacity(newSize); for (int i = originalSize; i < newSize; i++) list.add(0); |
| Rust  | a.resize(n \* 2, 0) |
| C++   | ↑ |
| C     | if (size >= capacity) { capacity *= 2; int \*temp = (int \*)realloc(dynamicArray, capacity \* sizeof(int)); if (temp == NULL) { free(dynamicArray); return 1; } dynamicArray = temp; dynamicArray[size++] = x; |

Slices Methods:

| Lang  | size | push |
| -- | -- | -- |
| JS    | arr.length | a.push(x) |
| Python| len(a) | a.append(x) |
| Go    | ↑ | a = append(a, x) |
| Rust  | vec.len() | a.push(x) |
| Java  | list.length | a.add(x) |
| C++   | vec.size() | a.push_back(x) |
| C     | - / size | a[size++] = x |

Heaps or Priority Queues:

| Lang  | pq |
| -- | -- |
| Python| heapq.heapify / heappush / heappop / heappushpop / heapreplace ... |
| Go    | (*h Heap) Len/Less/Swap/Push/Pop() ; heap.Init/Push/Pop(&h...) |
| Rust  | use std::collections::BinaryHeap; use std::cmp::Reverse; let mut pq = BinaryHeap::new(); pq.push((Reverse(0), node1)); let Some((Reverse(cost), cur)) = pq.pop(); |
| Java  | PriorityQueue<int[]> pq = new PriorityQueue<int[]>((a, b) -> a[0] - b[0]); pq.offer(new int[]{0, node1}); !pq.isEmpty(); int[] arr = pq.poll(); pq.offer(new int[]{cost + ncost, next}); |
| C++   | using pii = pair<int, int>; priority_queue<pii, vector\<pii\>, greater\<pii\>> pq; pq.emplace(0, node1); !pq.empty(); auto [cost, cur] = pq.top(); pq.pop(); pq.emplace(cost + ncost, next); |
| C     | typedef struct { Node \*arr; int capacity; int queueSize; cmp compare; } PriorityQueue; ... |

Stacks:

| Lang  | Stack |
| -- | -- |
| JS    | 通过Array |
| Python| 通过List |
| Go    | 通过Slice |
| Rust  | 通过Vec |
| Java  | Deque\<Integer\> stack = new LinkedList\<Integer\>(); stack.push(1); !stack.isEmpty(); stack.peek(); stack.pop(); / ArrayDeque |
| C++   | stack\<int\> stk; stk.push(1); !stk.empty(); stk.top(); stk.pop() |
| C     | int stk[n], top = 0; stk[top++] = 1; !top; top--; |

Deques & StringBuilders:

| Lang  | Deques & StringBuilders |
| -- | -- |
| JS    | a.unshift(x); a.push(x); q.reverse().join(''); |
| Python| q = Deque(), q.appendleft, q.append, q.popleft, q.pop<br>ans = "".join(q)<br>if head: ans = ans[::-1] |
| Go    | if head { q = append([]rune{ch}, q...) } else { q = append(q, ch) } if head { reverse(q) } return string(q) |
| Rust  | let mut q = Vec::\<char\>::new(); if head { q.insert(x); } else { q.push(x); } let ans: String = q.iter().rev().collect() |
| Java  | q.addFirst/Last(x); offer~; get~; peek~; remove~; poll~; StringBuilder ans = new StringBuilder(); while (!q.isEmpty()) ans.append(q.pollLast()); |
| C++   | deque\<char\> q; q.push_front(ch); q.push_back(ch); string ans = (head ? string{q.rbegin(), q.rend()} : string{q.begin(), q.end()}); |
| C     | char q1[len], q2[len]; int pos1 = 0, pos2 = 0; if (head) { q1[pos1++] = ch; } else { ... } char *ans = (char \*)malloc(sizeof(char) \* (len + 1)); int pos = 0; if (head) { for (int i = pos2 - 1; i >= 0; i--) ans[pos++] = q2[i]; memcpy(ans + pos, q1, sizeof(char) \* pos1); } else { ... } ans[pos1 + pos2] = '\0'; |

Linked List:

| Lang  | Deques & StringBuilders |
| -- | -- |
| Go | import 'containers/list' / list.List / e := l.Next(); e != nil / type entry struct{Key, Value} / et := e.Value.(entry) / e.Value =... / l.PushBack(entry{key, value}) / l.Remove(e) |
| Java | import java.util.LinkedList / LinkedList<X>/ Iterator<X> it = data[h].iterator(); / while (it.hasNext()) { X x = it.next() } / l.offerLast / l.remove(x) |
| C++ | list<pair<int, int>>; / for (auto it = data[h].begin(); it != data[h].end(); it++) / (*it).first / (*it).second =... / l.psuh_back(make_pair(key, value)) / l.erase(it) |
| C | struct List { int key; int val; struct List* next; }; ... |

Refs:
- 706. 设计哈希映射 https://leetcode.cn/problems/design-hashmap/solutions/
- 705. 设计哈希集合 https://leetcode.cn/problems/design-hashset/solutions/

## Math

Min / Max:

| Lang  | min(x, y) |
| -- | -- |
| JS    | Math.min |
| Python| min |
| Go    | - / min |
| Rust  | std::cmp::min |
| Java  | Math.min |
| C++   | min |
| C     | fmin |

Infinity / MaxValue:

| Lang  | Infinity | MaxInt | MaxValue |
| -- | -- | -- | -- |
| JS    | Infinity | Number.MAX_SAFE_INTEGER | Number.MaxValue |
| Python| inf / float('inf') | - | sys.float_info.max |
| Go    | -  | math.MaxInt8 / MaxInt16 / MaxInt32 / MaxInt64 | math.MaxFloat32 / MaxFloat64 |
| Rust  | std::f32::INFINITY / std::f64::INFINITY | i32::MAX / i64::MAX / max_value() | - |
| Java  | Double.POSITIVE_INFINITY / Double.NEGATIVE_INFINITY | Integer.MAX_VALUE / Long.MAX_VALUE | Float.MAX_VALUE / Double.MAX_VALUE |
| C++   | std::numeric_limits\<double\>::infinity() / std::numeric_limits\<float\>::infinity() | INT_MAX | FLT_MAX / DBL_MAX |
| C     | - | ↑ | ↑ |

## Control

io.read / write:

| Lang  | io.read / write |
| -- | -- | -- |
| Node.js | rl = readline.createInterface({ input, output }); rl.question('What do you think of Node.js? ', (answer) => {}); rl.close();
| Go    | fmt.Print("What do you want to do? ") / fmt.Scanln(&choice) |
| Shell | echo ... / read -p 'What do you want to do? ' choice |
}); |

switch case:

| Lang  | switch case |
| -- | -- | -- |
| Shell | case "$choice" in / '1') ... ;; |
| Go    | switch choice { / case "1": ... default: |

Refs:
- https://github.com/naltun/eyes.sh/blob/master/eyes.sh
- https://github.com/naltun/eyes/blob/master/main.go

## Misc

Sort:

| Lang  | Sort |
| -- | -- |
| JS    | intervals.sort((a, b) => a[0] - b[0]) |
| Python| intervals.sort(key=lambda x: x[0]) |
| Go    | sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] }) |
| Rust  | ranges.sort_by(\|a, b\| a[0].cmp(&b[0])) |
| Java  | Collections.sort / Arrays.sort(ranges, (a, b) -> a[0] - b[0]) / Arrays.sort(intervals, new Comparator<int[]>() { public int compare(int[] interval1, int[] interval2) { return interval1[0] - interval2[0]; } }); |
| C++   | sort(ranges.begin(), ranges.end()) |
| C     | int cmp(const void \*a, const void \*b) { return (\*(int \*\*)a)[0] - (\*(int \*\*)b)[0]; } qsort(ranges, rangesSize, sizeof(int \*), cmp) |

Const & Scientific notation:

| Lang  | Const |
| -- | -- |
| JS    | const MOD = 10**9 + 7 |
| Python| - / MOD = 10**9 + 7 |
| Go    | const mod = int(1e9 + 7) |
| Rust  | const MOD: i64 = 1_000_000_007 |
| Java  | static final int MOD = 1000000007 |
| C++   | static constexpr int mod = 1e9 + 7 / static const int base = 769 |
| C     | const int mod = 1e9 + 7 |

For Loop:

| Lang  | For i | For item | For i, item |
| -- | -- | -- | -- |
| JS    | for (let i = 0; i < list.length; i++) | for (let item of list) | - |
| Python| for i in range(0, len(list)) | for item in list | for i, v in enumerate(list) |
| Go    | for i := 0; i < len(list); i++ | - | for i, v := range list |
| Rust  | for i in 0..list.len() / for i in 1..=n | for item in list.iter() | for (i, item) in list.enumerate() / let indices = (0..list.len()).collect::<Vec<_>>(); for (index, item) in indices.iter().zip(list.iter()) |
| Java  | for (int i = 0; i < list.size(); i++) / IntStream.range(0, list.size()).forEach(i -> { String v = list.get(i) }) | for (Item item : list) | - |
| C++   | for (std::size_t i = 0; i < list.size(); ++i) | ↑ | - |
| C     | int a[] = {1, 2, 3, 4, 5}; int size = sizeof(a) / sizeof(a[0]); for (int i = 0; i < size; i++) | - | - |

Print:

| Lang  | Println | Printf |
| -- | -- | -- |
| JS    | console.log("x=%d", x) | - |
| Python| print("x=", x) | - |
| Go    | fmt.Println("x=", x) | fmt.Printf("x=%d\n", x) |
| Rust  | println!("x={}", x) | print!("x={}\n", x) |
| Java  | System.out.println("x="+x) | System.out.printf("x=%d\n", x) |
| C++   | std::cout << "x=" << x << std::endl | - |
| C     | - | printf("x=%d\n", x) |
