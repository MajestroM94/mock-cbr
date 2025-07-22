# Управление проектом
# Запуск сервера
run:
	go run cmd/server/main.go

# Очистка бд и временных файлов
clean:
	@if exist database\test.db ( \
		del /Q database\test.db && \
		echo Database was deleted \
	) else ( \
		echo File database\test.db not found \
	)

# Установка зависимостей
deps:
	go mod tidy

# Сваггер
swagger:
	swag init --generalInfo cmd/server/main.go --output ./docs


# Проверка кода
lint:
	go fmt ./...

# Навигатор команд
help:
	@echo "make run    - Server start"
	@echo "make clean  - Delete database and temporary files"
	@echo "make deps   - Installing dependencies"
	@echo "make swagger   - Creating swagger"
	@echo "make lint   - Code formatting"