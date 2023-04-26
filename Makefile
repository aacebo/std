fmt:
	gofmt -w .

test:
	go test -coverprofile cover.out

test.cov:
	go tool cover -html=cover.out
