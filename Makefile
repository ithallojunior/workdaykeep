name=bin/WorkDayKeep

build:
	@echo "Building..."
	go build -o $(name) -trimpath .

compressed:
	@echo "Building compressed version..."
	go build -o $(name) -ldflags="-s -w" -trimpath . 

run:
	go run .

linux64:
	@echo "Building compressed version for linux 64..."
	env GOOS=linux GOARCH=amd64 go build -o $(name) -ldflags="-s -w" -trimpath . 

windows:
	@echo "Building compressed version for windows..."
	env GOOS=windows GOARCH=amd64 go build -o $(name) -ldflags="-s -w" -trimpath . 