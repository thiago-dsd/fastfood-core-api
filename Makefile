build:
	echo "Generating docs"
	swag init --parseDependency
	echo "Compiling for your OS"
	go build -o ./bin/core-api-template

compile:
	echo "Generating docs"
	swag init --parseDependency
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o ./bin/core-api-template-linux-arm 
	GOOS=linux GOARCH=arm64 go build -o ./bin/core-api-template-linux-arm64 
	GOOS=windows GOARCH=386 go build -o ./bin/core-api-template-windows-386 
	GOOS=windows GOARCH=arm64 go build -o ./bin/core-api-template-windows-arm64

run:
	echo "Generating docs"
	swag init --parseDependency
	echo "Running your program"
	go run main.go

clean:
	echo "Cleaning outputs"
	@if [ -d "./bin" ]; then rm -r ./bin; fi
	swag fmt

setup: # Use this command to setup the environment

	echo "Setting docker container"
	docker run -d --name core-api-template-database -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres postgres:latest

stop-setup:
	docker stop core-api-template-database

remove-setup: # Use this command to remove all Docker containers and network
	docker rm core-api-template-database

init-setup: # Use this command to init the setup
	echo "Initializing setup"
	echo "Initializing postgres server"
	docker start core-api-template-database

