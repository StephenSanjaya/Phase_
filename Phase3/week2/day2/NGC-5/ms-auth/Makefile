.PHONY: proto
proto:
	protoc --go_out=. --go-grpc_out=. proto/*.proto
	@echo "Protoc file success!"

.PHONY: clean
clean:
	rm -r pb