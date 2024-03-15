```sh
# Go
go clean -testcache
go test -v ./...

# Python
pytest -v -s

# Rust
rustc hello_world.rs -o hello_world_exe
./hello_world_exe
rustc --test hello_world.rs -o hello_world_test_exe
./hello_world_test_exe
cargo test -- --nocapture

# Java
javac HelloWorld.java
java HelloWorld
mvn clean
mvn install
mvn test
mvn compile
java -cp target/classes me.fritx.App
mvn package
java -jar target/playground-1.0-SNAPSHOT.jar

# Solidity
npm install  # or pnpm install
npx hardhat test
```

Roadmap:
- [ ] benchmarks
- [ ] bigdata & spark & hadoop
- [x] solidity & hardhat
- [ ] c & c++ testing
- [x] mixed c & c++ with go
- [x] cargo test
- [x] rust & rustc --test
- [ ] gradle
- [x] java & maven & junit
- [ ] diagrams_bks
- [x] diagrams
- [ ] P2: python -m unittest discover -v
- [x] import cjk filenames & ~from sub-directories
- [x] pytest
- [x] go test

Refs:
- https://junit.org/junit5/docs/current/user-guide/#running-tests-build
