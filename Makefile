# Переменные
GO := go
GO_PKG := ./...

# Цель по умолчанию
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  test        - Run all tests"
	@echo "  coverage    - Run tests and generate HTML coverage report"
	@echo "  cover       - Alias for coverage"
	@echo "  lint        - Run golangci-lint"
	@echo "  all         - Run lint, tests and coverage"
	@echo "  help        - Show this help"

# Запуск всех тестов
.PHONY: test
test:
	$(GO) test -v $(GO_PKG)

# Генерация отчёта о покрытии в формате HTML
.PHONY: coverage cover
coverage cover:
	$(GO) test -coverprofile=coverage.out $(GO_PKG)
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: file://$(shell pwd)/coverage.html"

# Вывод покрытия в терминал (опционально)
.PHONY: cover-report
cover-report:
	$(GO) test -cover $(GO_PKG)

# Проверка кода с помощью golangci-lint
.PHONY: lint
lint:
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "❌ golangci-lint is not installed. Please install it:"; \
		echo "   https://golangci-lint.run/usage/install/"; \
		exit 1; \
	fi
	golangci-lint run

# Запуск всех проверок
.PHONY: all
all: lint test coverage