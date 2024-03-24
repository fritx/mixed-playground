
# Cheatsheet among languages

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

Static Arrays:

| Lang  | Array | copy | free |
| -- | -- | -- | -- |
| JS    | a = new Array(3) | - | - |
| Python| a = list() | - | - |
| Go    | var a [3]int = [3]int{} | copy(newArray, fixedArray) | - |
| Java  | int[] a = new int[3] | int[] newArray = new int[fixedArray.length * 2];<br>System.arraycopy(fixedArray, 0, newArray, 0, fixedArray.length); | - |
| C++   | int a[3]; | std::copy(std::begin(fixedArray), std::end(fixedArray), std::begin(newArray)) | delete[] a |
| C   | ↑ | - | free a |

Dynamic Arrays / Slices:

| Lang  | Slice | push |
| -- | -- | -- |
| JS    | 同Array | a.push(x) |
| Python| 同Array | a.append(x) |
| Go    | var a []int = []int{} | a = append(a, x) |
| Rust  | let a: Vec\<i32\> = vec![0; n] / Vec::new() | a.push(x) |
| Java  | List\<Integer\> a = new ArrayList\<Integer\>() / List\<String\> a = Arrays.asList("a", "b", "c") | a.add(x) |
| C++   | std::vector\<int\> a = {1, 2, 3} | a.push_back(x) |
| C     | char *list[] = {"a", "b", "c"} / int \*a = NULL; int size = 0; capacity = n; a = (int \*)malloc(capacity \* sizeof(int)); if (dynamicArray == NULL) return 1; | a[size++] = x |

Slices Realloc:

| Lang  | Realloc |
| -- | -- |
| Go  | newSlice := make([]int, 0, 2*cap(s))<br>newSlice = append(newSlice, s...) |
| Java  | list.ensureCapacity(newSize); for (int i = originalSize; i < newSize; i++) list.add(0); |
| Rust  | a.resize(n \* 2, 0) |
| C++   | ↑ |
| C     | if (size >= capacity) { capacity *= 2; int \*temp = (int \*)realloc(dynamicArray, capacity \* sizeof(int)); if (temp == NULL) { free(dynamicArray); return 1; } dynamicArray = temp; dynamicArray[size++] = x; |

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

Print

| Lang  | Println | Printf |
| -- | -- | -- |
| JS    | console.log("x=%d", x) | - |
| Python| print("x=", x) | - |
| Go    | fmt.Println("x=", x) | fmt.Printf("x=%d\n", x) |
| Rust  | println!("x={}", x) | print!("x={}\n", x) |
| Java  | System.out.println("x="+x) | System.out.printf("x=%d\n", x) |
| C++   | std::cout << "x=" << x << std::endl | - |
| C     | - | printf("x=%d\n", x) |
