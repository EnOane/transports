PROTO_DIR=./internal/grpc/proto
OUT_DIR=./pkg/grpc

PROTOC_GEN_GO=$(shell which protoc-gen-go)

all: generate

generate: gen-vd_engine

gen-vd_engine:
	mkdir -p $(OUT_DIR)/vd_engine
	protoc --go_out=$(OUT_DIR)/vd_engine --go_opt=paths=source_relative \
	       --go-grpc_out=$(OUT_DIR)/vd_engine --go-grpc_opt=paths=source_relative \
	       -I $(PROTO_DIR)/vd_engine $(PROTO_DIR)/vd_engine/*.proto

clean:
	rm -rf $(OUT_DIR)

.PHONY: all generate clean
