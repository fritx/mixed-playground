pub fn hello(name: &str) -> String {
    format!("Hello, {}!", name).to_string()
}

pub fn hello_2(name: &str) -> String {
    let mut owned_string: String = "Hello, ".to_owned();
    owned_string.push_str(name);
    owned_string.push_str("!");
    owned_string
}

#[test]
fn it_works() {
    assert_eq!(hello("world"), "Hello, world!");
    assert_eq!(hello_2("Alice"), "Hello, Alice!");
}
