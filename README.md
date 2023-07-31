# Golang

- Standard pacakges : /go/src
- GOROOT /bin /pkg
- GOPATH --> /bin  /pkg
- GOBIN --> 
- GOARCH --> targetted arch
- GOOS  --> target name of the OS

## create a go module

```
go mod init demo
```


 go run main.go
 go run --work main.go
 go build main.go
 ./main 
 go build -o hello main.go
 ./hello

- List of go target os and architectures
```
go tool dist list
 ```


                        
- To cross compile 
```
```
- to stripdown the binary 

```
go build -ldflags "-s -w" -o hello main.go
```

- to compile and link

```
go tool compile main.go

go tool link main.o
```