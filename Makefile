GO111MODULE=on
all:
	go build -o bin/h5fsServer ./main.go

doc:
	swag init

clean:
	go clean -i ./...
	@rm -rf bin/h5fsServer
