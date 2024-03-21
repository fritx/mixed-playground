use super::frequency_tracker_2671::FrequencyTracker;
use serde_json::Value;

#[test]
fn it_works() {
    let json_string = r#"
        [
            [["FrequencyTracker","add","add","hasFrequency"], [[],[3],[3],[2]], [null,null,null,true]],
            [["FrequencyTracker","add","deleteOne","hasFrequency"], [[],[1],[1],[1]], [null,null,null,false]],
            [["FrequencyTracker","hasFrequency","add","hasFrequency"], [[],[2],[3],[1]], [null,false,null,true]],
            [[], [], []]
        ]"#;
    let cases: Vec<Vec<Value>> = serde_json::from_str(json_string).unwrap();

    for case in cases.iter() {
        // let methods = &case[0].as_array().unwrap()[1..];
        let methods = &safe_slice(case[0].as_array().unwrap(), 1, None);
        let args = &safe_slice(case[1].as_array().unwrap(), 1, None);
        let want = &safe_slice(case[2].as_array().unwrap(), 1, None);
        run_test_case(methods, args, want);
    }
}

// fn run_test_case(methods: &Vec<Value>, args: &Vec<Value>, want: &Vec<Value>) {
fn run_test_case(methods: &[Value], args: &[Value], want: &[Value]) {
    let mut ft = FrequencyTracker::new();
    let n = methods.len();
    let mut ans: Vec<Value> = Vec::with_capacity(n);
    ans.resize(n, Value::Null);

    for (i, method) in methods.iter().enumerate() {
        let value: i32 = args[i][0].as_i64().unwrap() as i32;
        match method.as_str().unwrap() {
            "add" => {
                ans[i] = Value::from(ft.add(value));
            }
            "deleteOne" => {
                ans[i] = Value::from(ft.delete_one(value));
            }
            "hasFrequency" => {
                ans[i] = Value::from(ft.has_frequency(value));
            }
            _ => panic!("Unknown method: {}", method),
        }
    }
    assert_eq!(&ans, want);
}

// @author 文心一言
fn safe_slice<T>(slice: &[T], start: usize, end: Option<usize>) -> &[T] {
    // 确保start不会超过slice的长度，并且end不会超过slice的起始位置
    let start = start.min(slice.len());
    let end = end
        .unwrap_or_else(|| slice.len())
        .min(slice.len())
        .max(start); // 确保end不小于start

    // 返回从start到end的切片，如果end小于start，则返回空切片
    &slice[start..end]
}
// @author BingChat
fn slice<'a, T>(arr: &'a [T], start: usize, end: usize) -> &'a [T] {
    let start = std::cmp::min(start, arr.len());
    let end = std::cmp::min(end, arr.len());
    &arr[start..end]
}
