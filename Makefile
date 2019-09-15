export GO111MODULE=on

.PHONY: clean build test deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/awslambda/getone cmd/awslambda/getone/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	./node_modules/.bin/sls deploy --verbose

test:
	go test -v ./...
