use std::collections::HashSet;

pub fn assert_arrays_equal_unordered(a: &[i32], b: &[i32]) {
    let set_a: HashSet<i32> = a.iter().cloned().collect();
    let set_b: HashSet<i32> = b.iter().cloned().collect();
    assert_eq!(set_a, set_b);
}

// fn assert_arrays_equal_unordered_2(a: &[i32], b: &[i32]) {
//     let mut sorted_a: Vec<i32> = a.to_vec();
//     let mut sorted_b: Vec<i32> = b.to_vec();
//     sorted_a.sort();
//     sorted_b.sort();
//     assert_eq!(sorted_a, sorted_b);
// }
