build: 
	@go build -o ./bin/app
	@cd ./www/svelte && npm run build

run: build
	@./bin/app

clean:
	@rm -rf ./bin
