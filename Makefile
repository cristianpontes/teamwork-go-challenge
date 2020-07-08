SRC = $$(go list  -f {{.Dir}} ./... | grep -vE "./*/testing" | grep -vE "./*/gen" | grep -v /vendor/)
PACKAGES = $$(go list ./... | grep -vE "./*/testing" | grep -vE "./*/gen" | grep -v /vendor/)

config:
	go install https://github.com/dave/courtney

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
	courtney -e $(PACKAGES)

coverage-report:
	courtney -e $(PACKAGES)
	go tool cover -html=coverage.out