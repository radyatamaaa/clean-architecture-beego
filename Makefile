BINARY=engine
BRANCH=$(shell git rev-parse --abbrev-ref 'HEAD')
commit:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
	git commit

test: 
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

unittest:
	go test -short  ./...

swagger_documentation:
	swag init -g main.go --output swagger

generate_protoc:
    protoc -I . --go_out=. --go_opt paths=source_relative --go-grpc_out=.  --go-grpc_opt paths=source_relative ./product_service.proto

generate_protoc_gateway:
    protoc -I . --grpc-gateway_out=. --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true ./product_service.proto

generate_import_package_grpc_gateway_api:
    curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > annotations.proto
    curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > http.proto