generate:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/grpctask.proto

build:
	go build cmd/app/app.go

run:
    docker run -dp 50051:50051 pointservice