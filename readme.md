# HTTP Calculator

Мой первый проект :)

Это итоговая задача 1 модуля в Яндекс Лицее. Нужно было создать веб-сервис для вычисления математических выражений. Проект достаточно простой - отправляете выражение через POST-запрос в формате JSON, а в ответ получаете json ответ - ответ на выражение или ошибка:

```
#запрос
{"expression": "2+2*2"}

#ответ
{"result": 6}
```
### Что умеет калькулятор?
- Все базовые математические операции (+, -, *, /)
- Правильно обрабатывает приоритет операций (сначала умножение/деление, потом сложение/вычитание)
- Поддерживает скобки для изменения приоритета - можно написать (2+2)*2
- Работает с десятичными числами (например, 2.5 + 3.7)
- Возвращает понятные ошибки, если что-то пошло не так

*КАЛЬКУЛЯТОР НЕ УМЕЕТ РАБОТАТЬ С ОТРИЦАТЕЛЬНЫМИ ЧИСЛАМИ*


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

(очень надеюсь что у вас будут работать курлы)

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
│   └── handler/        # HTTP обработчик
├── go.mod
└── README.md
```

### чут чут уточнений для проверяющего
- я сделал handler_test, тк в таком случае я сразу тестирую и калькулятор. не думаю что это большая проблема, ведь в мейне особо и нечего тестировать.

- валидацию порта не делал, там и так если вы неккоректный порт ввели понятную ошибку кинет. При этом сервер не запускался, значит ничего выдавать не должен Ж)

- покрытие тестами у меня вышло на 75% 

- мой юз в тг ```@zxcbilka```

спасибо за проверку, всего хорошего!


![эщкерее](internal/important.jpg)
