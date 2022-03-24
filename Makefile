build:
	@cd sumr && go build

fmt:
	@cd sumr && go fmt

install:
	@cd sumr && go install

test:
	@cd sumr && go test

clean:
	@rm sumr/sumr > /dev/null 2>&1 || true
