test:
	go install util
	go install util/file
	go install util/json
	go test

clean:
	go clean util
	go clean util/file
	go clean util/json
