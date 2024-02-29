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
```

Roadmap:
- [ ] benchmarks
- [x] cargo test
- [x] rust & rustc --test
- [ ] maven
- [ ] java & junit
- [x] mingrammer/diagrams
- [ ] P2: python -m unittest discover -v
- [x] pytest
- [x] go test
