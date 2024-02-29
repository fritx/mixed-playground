pub fn hello(name: &str) -> String {
    format!("Hello, {}!", name).to_string()
}

#[test]
fn it_works() {
    assert_eq!(hello("world"), "Hello, world!");
    assert_eq!(hello("Alice"), "Hello, Alice!");
}
