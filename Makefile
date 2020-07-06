SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "*/gen/*")

config:
	go mod vendor

	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.28.0

	if [ -e ./.git/hooks/pre-commit ]; then\
		rm ./.git/hooks/pre-commit;\
	fi
	cp ./git/pre-commit.sh ./.git/hooks/pre-commit
	chmod u+x .git/hooks/pre-commit

format:
	@gofmt  -l -w $(SRC)
	@goimports -local github.com/cristianpontes/teamwork-go-challange -l -w $(SRC)

lint:
	./bin/golangci-lint run --config=./git/golangci.yml

install:
	go build -o $(GOPATH)/bin/tw-go-challenge ./cmd/customer

test:
	go test ./...

coverage-report:
	 go test -coverprofile=coverage.out ./...
	 go tool cover -html=coverage.out