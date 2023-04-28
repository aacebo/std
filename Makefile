fmt:
	gofmt -w .

test:
	go clean -testcache
	go test ./... -coverprofile cover.out

test.cov: test
	go tool cover -html=cover.out
