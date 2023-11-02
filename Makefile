env:
	@cp .env_template  .env

protoc-order:
	@protoc --go_out=. --go-grpc_out=. internal/adapters/grpc/protofiles/order.proto

grpc-evans:
	@evans -r repl

graphql-generate:
	@go run github.com/99designs/gqlgen generate
