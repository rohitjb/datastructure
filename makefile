COVERFILE=coverage.out
COVERAGE=78.00

run:
	go run main.go

test:
	go test ./... -tags=unittest -covermode=atomic -coverprofile=$(COVERFILE) -v

cover-visual: test
	@echo 'Check coverage in your browser'
	@go tool cover -html=$(COVERFILE)

install:
	go install ./...

clean:
	go clean ./...
	rm -f transactions coverage.out
	rm -rf ./dist

vendor:
	go mod tidy
	go mod vendor