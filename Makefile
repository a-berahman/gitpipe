start:
	@go mod init 
	@go mod tidy

run:
	go run main.go -c <*.yaml dir> serve