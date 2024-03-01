```sh
# Go
go clean -testcache && go test -v ./...

# Python
pytest -v -s

# Rust
rustc hello_world.rs -o hello_world_exe
./hello_world_exe
rustc --test hello_world.rs -o hello_world_test_exe
./hello_world_test_exe
cargo test

# Java
javac HelloWorld.java
java HelloWorld
mvn install
mvn test
mvn compile
java -cp target/classes me.fritx.App
mvn package
java -jar target/playground-1.0-SNAPSHOT.jar
```

Roadmap:
- [ ] benchmarks
- [ ] c & c++
- [x] cargo test
- [x] rust & rustc --test
- [ ] P2: java tests by jar
- [ ] gradle
- [x] java & maven & junit
- [x] mingrammer/diagrams
- [ ] P2: python -m unittest discover -v
- [x] pytest
- [x] go test

Refs:
- https://junit.org/junit5/docs/current/user-guide/#running-tests-build
