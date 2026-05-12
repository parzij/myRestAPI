# Простой REST API с использованием Echo

Это базовый проект REST API на языке Go с использованием веб-фреймворка **Echo**. API поддерживает простые операции CRUD (Create, Read, Update, Delete) с данными в формате JSON и включает Swagger-документацию.

## 🚀 Технологии

- **Golang** - язык программирования
- **Echo v4** - веб-фреймворк
- **Swaggo/Swag** - генерация Swagger документации
- **Swagger UI** - интерактивная документация API
- **Godotenv** - загрузка переменных окружения
- **JSON** - формат обмена данными

## 📁 Структура проекта

```
MYRESTAPI
├── cmd/api
│   └── main.go              # Точка входа в приложение
├── docs
│   ├── docs.go              # Автогенерируемый Swagger
│   ├── swagger.json         
│   └── swagger.yaml         
├── internal
│   ├── handler
│   │   └── handler.go       # Обработчики HTTP запросов
│   ├── models
│   │   └── models.go        # Модели данных
│   └── router
│       └── router.go        # Настройка маршрутов
├── .env                     
├── .gitignore               
├── go.mod                   
├── go.sum                   
├── LICENSE                  
└── README.md                
```

## 🔌 Доступные эндпоинты

### 📥 Получение всех сообщений

```http
GET http://localhost:port/messages
```

**Ответ (200 OK):**
```json
[
  {
    "text": "Первое сообщение"
  },
  {
    "text": "Второе сообщение"
  }
]
```

**Ответ (404 Not Found):**
```json
{
  "status": "Error",
  "message": "No messages found"
}
```

---

### ➕ Создание сообщения

```http
POST http://localhost:port/messages
Content-Type: application/json

{
  "text": "Новое сообщение"
}
```

**Ответ (200 OK):**
```json
{
  "status": "Success",
  "message": "Message added successfully"
}
```

**Ответ (400 Bad Request):**
```json
{
  "status": "Error",
  "message": "Could not add the message"
}
```

---

### ✏️ Обновление сообщения

```http
PUT http://localhost:port/messages/0
Content-Type: application/json

{
  "text": "Обновлённый текст"
}
```

Где `0` - индекс сообщения в массиве.

**Ответ (200 OK):**
```json
{
  "status": "Success",
  "message": "Message updated successfully"
}
```

**Ответ (400 Bad Request):**
```json
{
  "status": "Error",
  "message": "Invalid message ID"
}
```

---

### 🗑️ Удаление сообщения

```http
DELETE http://localhost:port/messages/0
```

Где `0` - индекс сообщения для удаления.

**Ответ (200 OK):**
```json
{
  "status": "Success",
  "message": "Message deleted successfully"
}
```

**Ответ (400 Bad Request):**
```json
{
  "status": "Error",
  "message": "Invalid message ID"
}
```

## 🏗️ Архитектура проекта

### Модульная структура

Проект разделён на логические модули:

- **`cmd/api/main.go`** - точка входа, инициализация сервера и роутера
- **`internal/handler/handler.go`** - обработчики HTTP запросов (бизнес-логика)
- **`internal/models/models.go`** - структуры данных и глобальные переменные
- **`internal/router/router.go`** - настройка маршрутизации
- **`docs/`** - автогенерируемая Swagger документация

### Модели данных

**Message** - структура сообщения:
```go
type Message struct {
    Text string `json:"text"`
}
```

**Response** - структура ответа API:
```go
type Response struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}
```
##  Лицензия

Проект открыт для использования и модификаций. См. файл [LICENSE](LICENSE) для подробностей.
