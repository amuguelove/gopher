build:
# 一定要注意 Makefile 中的缩进，否则 make build 可能报错 Nothing to be done for build
# protoc 命令前边是一个 Tab，不是四个或八个空格
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/gopher/shippy/feature1/consignment-service proto/consignment/consignment.proto