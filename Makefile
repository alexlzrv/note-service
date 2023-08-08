LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=5434 dbname=note-service user=note-service-user password=note-service-password sslmode=disable"

.PHONY: generate
generate:
					mkdir -p pkg/note_v1
					protoc 	--proto_path api/note_v1	\
                    			--go_out=pkg/note_v1 --go_opt=paths=source_relative \
                    			--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
                    			api/note_v1/note.proto

.PHONY: docker-run
docker-run:
	docker-compose up --build -d

.PHONY: docker-stop
docker-stop:
	docker-compose down

.PHONY: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

.PHONY: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v
