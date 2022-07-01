check-linux:
	@consul version || echo "Please Open this link 'https://learn.hashicorp.com/tutorials/consul/get-started-install#install-consul'"
check-osx:
	@consul version || (brew tap hashicorp/tap && brew install hashicorp/tap/consul)
docker-up:
	@docker-compose up -d
docker-down:
	@docker-compose down
populate:
	@consul kv put arithmetic/features/addition true
	@consul kv put arithmetic/features/substract false
	@consul kv put arithmetic/features/divide false
	@consul kv put arithmetic/features/multiply true
go:
	@go run main.go
run-all:
	@echo "Well this command not configured yet 😜"