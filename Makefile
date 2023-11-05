test:
	go test -v ./...
cover:
	go test ./... -v -covermode=atomic -coverprofile=coverage.out

cover-html:
	make cover
	go tool cover -html=coverage.out

