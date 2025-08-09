run:
	@env $$(cat .env.local | xargs) go run cmd/main.go
