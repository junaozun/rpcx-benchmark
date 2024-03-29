win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/server.exe server.go

mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/server server.go

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/server server.go
