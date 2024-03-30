docker run --rm -it --name golang -v $(pwd)/:/go/app golang:1.22 bash

apt update
sudo apt install -y protobuf-compiler
protoc --version # Ensure compiler version is 3+

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# toda vez que alterar esse arquivo rode isso no terminal
protoc --go_out=. --go-grpc_out=. proto/course_category.proto

go run cmd/server/main.go

go install github.com/ktr0731/evans@latest
evans --version

evans -r repl 
#ou
evans -r repl --proto /home/thales/projetos/grpc-golang/proto/course_category.proto

package pb
show service

show package
show CategoryService
service CategoryService
call CreateCategory

