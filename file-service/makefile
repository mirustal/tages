gen-grpc:
	rm -f pkg/pb/*.go
	protoc  --proto_path=./api/file \
		--proto_path=./api/file \
		--go_out=pkg/pb \
		--go_opt=paths=source_relative \
			--go-grpc_out=pkg/pb \
		--go-grpc_opt=paths=source_relative  \
			api/file/*.proto api/file/*.proto