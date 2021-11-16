set CGO_ENABLED=0
set GOOS=darwin
set GOARCH=amd64
go build -gcflags=-m -ldflags="-w -s" -o tmp/mrng.darwin main.go static_handle.go
upx tmp/mrng.darwin
