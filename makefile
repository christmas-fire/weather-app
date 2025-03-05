build:
	@echo "Building application..."
	@go build -o weather cmd/main.go

run:
	@go build -o weather cmd/main.go
	@./weather

clean:
	@echo "Cleaning executable files.."
	@rm -rf weather