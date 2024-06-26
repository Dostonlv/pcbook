gen:
	protoc --proto_path=proto --go_out=pb --go-grpc_out=pb proto/*.proto   

clean:
	rm pb/*.go

server:
	go run cmd/server/main.go -port 8080
client:
	go run cmd/client/main.go -address 0.0.0.0:8080
test:
	@echo "\033[92mTest starting\033[0m"
	go test -cover  -race -v  ./...

cert:
	cd cert; ./gen.sh; cd ..

.PHONY: gen clean server client test cert