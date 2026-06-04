PROTO_DIR := proto
PROTO_SRC := $(wildcard $(PROTO_DIR)/*.proto)
GO_OUT := .
MIGRATE_PATH=services/profile-service/cmd/migrations

.PHONY: generate-proto
generate-proto:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GO_OUT) \
		--go-grpc_out=$(GO_OUT) \
		$(PROTO_SRC)


.PHONY: migration
migration:
	@read -p "Enter service name: " service; \
	read -p "Enter migration name: " name; \
	migrate create -ext sql -dir services/$$service/cmd/migrations -seq $$name