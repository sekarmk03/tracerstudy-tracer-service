gen:
	protoc --proto_path=proto --go_out=./pb --go-grpc_out=./pb proto/*.proto

.PHONY:
	gen