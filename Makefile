PROTO_DIR=./internal/grpc/proto
OUT_DIR=./pkg/grpc

PROTOC_GEN_GO=$(shell which protoc-gen-go)

all: generate

generate: gen-dld gen-tg

gen-dld:
	mkdir -p $(OUT_DIR)/dld
	protoc --go_out=$(OUT_DIR)/dld --go_opt=paths=source_relative \
	       --go-grpc_out=$(OUT_DIR)/dld --go-grpc_opt=paths=source_relative \
	       -I $(PROTO_DIR)/cli_download $(PROTO_DIR)/cli_download/*.proto

gen-tg:
	mkdir -p $(OUT_DIR)/tg
	protoc --go_out=$(OUT_DIR)/tg --go_opt=paths=source_relative \
	       --go-grpc_out=$(OUT_DIR)/tg --go-grpc_opt=paths=source_relative \
	       -I $(PROTO_DIR)/tg_client $(PROTO_DIR)/tg_client/*.proto

clean:
	rm -rf $(OUT_DIR)

.PHONY: all generate clean
