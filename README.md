# go-example
go-example

## ftp server
build
- go build -ldflags "-s -w" cmd/ftp/ftp.go
- go run cmd/ftp/ftp.go


## sock5 server
build
- go build -ldflags "-s -w" cmd/socks5/socks5.go
- go run cmd/socks5/socks5.go