#![allow(dead_code)]
mod costs_of_paths_in_bt_2673;
mod dp;
mod graphs;
mod hello_module;
mod strings;
use hello_module::hello;

fn main() {
    println!("{}", hello("main.rs"));
}

#[test]
fn it_works() {
    assert_eq!(1 + 1, 2);
}
