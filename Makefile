### Code validation
check: ## Run all checks: test, lint, tidy
	@bash scripts/check.sh

test: ## Run tests for all go packages
	@bash scripts/test.sh

build:
	@bash scripts/build.sh

start:
	@bash scripts/start.sh

clean:
	@bash scripts/cleanup.sh