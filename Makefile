.DEFAULT_GOAL = test
.PHONY: FORCE build build_race clean test test_race test_linters module format print_ast

export GOPROXY = https://goproxy.cn,direct

# Build

build: go-readability
	go build -v

build_race:
	go build -race -v

clean:
	rm -f go-readability
	rm -f test/path
	rm -f tools/Dracula.itermcolors
	rm -f tools/goreleaser
	rm -f tools/svg-term
	rm -rf tools/node_modules

# Test
test: export GO_READABILITY = true
test: build
	GO_READABILITY_TEST_RUN=1 ./go-readability run -v
	GO_READABILITY_TEST_RUN=1 go test -v -parallel 2 ./...

test_race: build_race
	GO_READABILITY_TEST_RUN=1 ./go-readability run -v --timeout=5m

test_linters:
	GO_READABILITY_TEST_RUN=1 go test -v ./test -count 1 -run TestSourcesFromTestdataWithIssuesDir/$T

module:
	go mod tidy
	go mod verify

print_ast:
	@pushd tools/print_ast &> /dev/null ; go install -v ; popd &> /dev/null
