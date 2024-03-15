```sh
# Go
go clean -testcache
go test -v ./...
# w/ coverage
go test -cover ./...
go test -coverprofile=go_coverage.out ./...
go tool cover -func=go_coverage.out
go tool cover -html=go_coverage.out  # serving html

# Python
pytest -v -s
# w/ coverage
pip install pytest-cov
pytest --cov=.
pytest --cov=. --cov-report html:py_coverage
# html report is under `py_coverage` dir

# Rust
rustc hello_world.rs -o hello_world_exe
./hello_world_exe
rustc --test hello_world.rs -o hello_world_test_exe
./hello_world_test_exe
cargo build
cargo build --release
cargo test -- --nocapture
cargo test --release -- --nocapture

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
REPORT_GAS=true npx hardhat test
npx hardhat coverage
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
- [x] hacking cjk filenames & sub-directories & cwd
- [x] pytest
- [x] go test

Refs:
- https://doc.rust-lang.org/rustc/instrument-coverage.html
- https://pytest-cov.readthedocs.io/en/latest/reporting.html
- https://github.com/golang/vscode-go/blob/master/docs/features.md#code-coverage
- https://brantou.github.io/2017/05/24/go-cover-story/
- https://junit.org/junit5/docs/current/user-guide/#running-tests-build
