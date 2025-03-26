PROTO_DIR=./internal/proto
OUT_DIR=./pkg

PROTOC_GEN_GO=$(shell which protoc-gen-go)

all: generate

generate: gen-dld gen-tg

gen-dld:
	mkdir -p $(OUT_DIR)/dld
	protoc --go_out=$(OUT_DIR)/dld --go_opt=paths=source_relative \
	       --go-grpc_out=$(OUT_DIR)/dld --go-grpc_opt=paths=source_relative \
	       -I $(PROTO_DIR) $(PROTO_DIR)/*.proto

gen-tg:
	mkdir -p $(OUT_DIR)/tg
	protoc --go_out=$(OUT_DIR)/tg --go_opt=paths=source_relative \
	       --go-grpc_out=$(OUT_DIR)/tg --go-grpc_opt=paths=source_relative \
	       -I $(PROTO_DIR) $(PROTO_DIR)/*.proto

clean:
	rm -rf $(OUT_DIR)/dld
	rm -rf $(OUT_DIR)/tg

.PHONY: all generate clean
