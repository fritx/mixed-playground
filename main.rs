#![allow(dead_code)]
mod data_structures;
mod dp;
mod graphs;
mod hello_module;
mod math_ops;
mod pq;
mod strings;
mod trees;
mod utils;
use hello_module::hello;

fn main() {
    println!("{}", hello("main.rs"));
}

#[test]
fn it_works() {
    assert_eq!(1 + 1, 2);
}
