default: help

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

run: # 
	@go run cmd/main.go 

test: # Run teste with coverage
	@go clean -testcache
	@go test -v ./... -coverprofile=coverage.out
	@go tool cover -func coverage.out | awk 'END{print sprintf("coverage: %s", $$3)}'

test-cover: test # Run test and open cover at html
	go tool cover -html=coverage.out 