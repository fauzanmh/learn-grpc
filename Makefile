gen:
	protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto
clean:
	rm pb/*.go
run:
	go run main.go
test:
	go test -cover -race ./...

grpc-request-gen: 
	protoc proto/*.proto --go_out=plugins=grpc:pb

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080