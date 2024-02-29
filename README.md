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
rustc --test costs_of_paths_in_bt_2673.rs -o costs_of_paths_in_bt_2673_exe
./costs_of_paths_in_bt_2673_exe
```

Roadmap:
- [ ] benchmarks
- [ ] cargo test
- [x] rust & rustc --test
- [ ] maven
- [ ] java & junit
- [ ] P2: python -m unittest discover -v
- [x] pytest
- [x] go test
