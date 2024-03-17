#![allow(dead_code)]
mod dp;
mod graphs;
mod hello_module;
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
