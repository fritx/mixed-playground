mod hello_module;
use hello_module::hello;

fn main() {
    println!("{}", hello("Alice"));
}

#[test]
fn it_works() {
    assert_eq!(hello("world"), "Hello, world!");
    assert_eq!(hello("Alice"), "Hello, Alice!");
}
