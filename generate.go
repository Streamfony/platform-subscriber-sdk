package main

//go:generate protoc -I ./specs --go_out=./subscriber --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_out=./subscriber --doc_out=./docs --doc_opt=markdown,api.md ./specs/subscriber.proto
