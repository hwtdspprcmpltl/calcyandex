# HTTP Calculator

Веб-сервис для вычисления арифметических выражений. Поддерживает базовые арифметические операции (+, -, *, /) и работу со скобками.

## Функциональность

Сервис предоставляет endpoint для вычисления математических выражений:
- Поддерживаемые операции: +, -, *, /
- Поддержка скобок для изменения приоритета операций
- Поддержка десятичных чисел

## Установка

### Предварительные требования
- Go 1.19 или выше
- Git

### Установка проекта
```bash
# клонируем репозиторий:)
git clone https://github.com/hwtdspprcmpltl/calcyandex
```
```bash
# переходим в папку с проектом
cd calcyandex
```

## Запуск сервера
```bash
# пример запуска с настройками по умолчанию (порт 8080)
go run cmd/main.go
```
```bash
# пример запуска с указанием порта
go run cmd/main.go -port 3000
```

## Примеры Использования

### Endpoint: `/api/v1/calculate`

#### 1. Успешное вычисление
```bash
curl -i --location http://localhost:8080/api/v1/calculate --header "Content-Type: application/json" --data "{\"expression\": \"(2+2)*2\"}"
```
##### Ответ (200 OK)
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 21 Dec 2024 12:14:35 GMT
Content-Length: 13

{"result":8}
```

#### 2. Ошибка в расстановке скобок
```bash
curl -i --location http://localhost:8080/api/v1/calculate --header "Content-Type: application/json" --data "{\"expression\": \"(2+2)*2)\"}"
```
##### Ответ (422 Unprocessable Entity)
```bash
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
Date: Sat, 21 Dec 2024 12:12:16 GMT
Content-Length: 66

{"error":"ошибка в расставлении скобок"}
```

#### 3. Деление на ноль
```bash
curl -i --location http://localhost:8080/api/v1/calculate --header "Content-Type: application/json" --data "{\"expression\": \"2/0\"}"
```
##### Ответ (422 Unprocessable Entity)
```bash
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
Date: Sat, 21 Dec 2024 12:15:20 GMT
Content-Length: 45

{"error":"деление на ноль"}
```

#### 4. Неверный метод запроса
```bash
curl -i --location --request GET http://localhost:8080/api/v1/calculate
```
##### Ответ (405 Method Not Allowed)
```bash
HTTP/1.1 405 Method Not Allowed
Content-Type: application/json
Date: Sat, 21 Dec 2024 12:16:45 GMT
Content-Length: 43

{"error":"не пост запрос"}
```

## Коды ответов
- 200: Успешное вычисление
- 422: Ошибка в выражении (неверный формат, деление на ноль и т.д.)
- 500: Внутренняя ошибка сервера
- 405: Неверный метод запроса (только POST разрешен)

## Запуск тестов
```bash
# запуск всех тестов
go test ./...

# запуск тестов с информации о покрытии
go test -cover ./...
```


## Структура проекта
```
.
|── cmd/                 #Точка входа
├── internal/
│   ├── calc/           # Логика калькулятора
│   └── handler/        # HTTP обработчики
├── go.mod
└── README.md
```

# ps

