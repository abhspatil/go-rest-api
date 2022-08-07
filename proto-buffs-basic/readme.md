Install protobuf for golang

go get github.com/golang/protobuf
go get github.com/golang/protobuf/proto

Generate Proto files
protoc -I=protos/ --go_out=pbout/ protos/users.proto
or
protoc --proto_path=protos protos/users.proto --go_out=pbout/