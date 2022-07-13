test:
	@echo "running tests and coverage..."
	@mkdir -p tests-output/coverage
	@go test -p 1 -mod=vendor ./... -v -coverprofile tests-output/cover.out
	@go tool cover -func tests-output/cover.out
	@go tool cover -html tests-output/cover.out