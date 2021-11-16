set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -gcflags=-m -ldflags="-w -s" -o tmp/mrng.exe main.go static_handle.go
upx tmp/mrng.exe
