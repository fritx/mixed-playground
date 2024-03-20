struct Solution;

impl Solution {
    pub fn min_non_zero_product(p: i32) -> i32 {
        fn fast_pow(mut x: i64, mut n: i64, mod_val: i64) -> i64 {
            let mod_val = mod_val as i64;
            let mut res = 1;
            while n != 0 {
                if n & 1 == 1 {
                    res = (res * x) % mod_val;
                }
                x = (x * x) % mod_val;
                n >>= 1;
            }
            res
        }

        if p == 1 {
            return 1;
        }
        let mod_val = 1_000_000_007;
        let x = fast_pow(2, p as i64, mod_val) - 1;
        let y = 1i64 << (p - 1);
        (fast_pow(x - 1, y - 1, mod_val) * x % mod_val) as i32
    }
}

#[test]
fn it_works() {
    let json = r#"[
            [1, 1],
            [2, 6],
            [3, 1512],
            [31, 138191773]
        ]"#;
    let cases: Vec<Vec<i32>> = serde_json::from_str(json).unwrap();

    for case in cases {
        let p = case[0];
        let want = case[1];
        let ans = Solution::min_non_zero_product(p);
        assert_eq!(ans, want);
    }
}
