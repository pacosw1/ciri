# ciri
A lightweight programming language for iot devices



##Making changes to goyacc
Edit ``src/compiler/parser.y``

and run
```go run golang.org/x/tools/cmd/goyacc -l -o generated.go parser.y```



##Testing

Run ``go test ./...``