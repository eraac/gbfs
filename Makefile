run:
	go run example/*.go

code-quality:
	@echo "== GOLINT =="
	@find . -type d | xargs -L 1 golint
	@echo "== GO VET =="
	@go tool vet .
