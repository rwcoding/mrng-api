set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -gcflags=-m -ldflags="-w -s" -o tmp/mrng.linux main.go static_handle.go
upx tmp/mrng.linux