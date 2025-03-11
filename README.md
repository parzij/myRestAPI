# Простой REST API с использованием Echo (Golang)

Это базовый проект REST API на языке Go с использованием веб-фреймворка **Echo**. API поддерживает простые операции CRUD (создание, чтение, обновление, удаление) с данными в формате JSON.

## Технологии

- Golang
- Echo framework
- JSON

## Структура проекта

- `main.go` — основной файл с логикой API.
- Используем фреймворк Echo для обработки HTTP-запросов.

## Установка и запуск

1. Убедитесь, что у вас установлен Go.

2. Склонируйте репозиторий:

```bash
git clone <url вашего репозитория>
cd <ваша папка>
```

3. Установите зависимости:

```bash
go mod init echo_api
go get github.com/labstack/echo/v4
go get github.com/labstack/gommon/log
go mod tidy
```

3. Запустите сервер:

```bash
go run main.go
```

API будет доступен на `http://localhost:8000`.

## Доступные запросы

### Получение списка сообщений

```http
GET http://localhost:8000/messages
```

### Добавление нового сообщения

```http
POST http://localhost:8000/messages
Content-Type: application/json

{
  "text": "Ваше сообщение"
}
```

### Обновление существующего сообщения

```http
PUT http://localhost:8000/messages/{id}
Content-Type: application/json

{
  "text": "Обновлённый текст сообщения"
}
```

Где `{id}` – это индекс сообщения в списке (нумерация с 0).

### Удаление сообщения

```http
DELETE http://localhost:8000/messages/{id}
```

Где `{id}` – индекс сообщения, которое требуется удалить.

## Структура кода

Проект имеет следующие основные компоненты:

- Структура данных сообщения:

```go
type Message struct {
	Text string `json:"text"`
}
```

- Структура ответа сервера:

```go
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
```

- Примеры обработчиков запросов:

#### GET-запрос

```go
func Gethandler(c echo.Context) error {
	if len(messages) == 0 {
		return c.JSON(http.StatusNotFound, Response{
		Status:  "Error",
		Message: "No messages found",
	})
}
	return c.JSON(http.StatusOK, messages)
}
```

- POST-запрос добавляет сообщение в список:

```go
func PostHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the message",
	})
}
	messages = append(messages, message)
	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message added successfully",
	})
}
```

## Обработка ошибок

Все методы возвращают JSON-ответы с соответствующими статусами HTTP и понятными сообщениями об ошибках.

## Использование в веб-разработке

Данный API можно использовать как backend для приложений:
- Реализуйте клиентский интерфейс на React или другом JavaScript-фреймворке.
- Используйте API для интеграции с мобильными приложениями.
- Можно расширить функционал, подключив базу данных для хранения данных вместо временного массива.

## Лицензия

Проект открыт для использования и модификаций.

