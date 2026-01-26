

# REST API - Подробное руководство

## Что такое REST?

**REST** (Representational State Transfer) — архитектурный стиль для проектирования веб-сервисов. REST API — это способ взаимодействия между клиентом и сервером через HTTP протокол.

## Основные принципы REST

### 1. **Ресурсы (Resources)**
Все данные представлены как ресурсы, которые идентифицируются через URL.

```
/users          - коллекция пользователей
/users/1         - конкретный пользователь с ID=1
/events          - коллекция событий
/events/5        - конкретное событие с ID=5
```

### 2. **HTTP Методы (Verbs)**
Каждый HTTP метод имеет определенное назначение:

- **GET** - получение данных (чтение)
- **POST** - создание нового ресурса
- **PUT** - полное обновление ресурса
- **PATCH** - частичное обновление ресурса
- **DELETE** - удаление ресурса

### 3. **Статус коды HTTP**
Сервер должен возвращать правильные HTTP статус коды:

**Успешные ответы (2xx):**
- `200 OK` - успешный запрос
- `201 Created` - ресурс успешно создан
- `204 No Content` - успешный запрос без содержимого

**Ошибки клиента (4xx):**
- `400 Bad Request` - неверный запрос
- `401 Unauthorized` - требуется аутентификация
- `403 Forbidden` - доступ запрещен
- `404 Not Found` - ресурс не найден
- `405 Method Not Allowed` - метод не разрешен
- `409 Conflict` - конфликт (например, дубликат)

**Ошибки сервера (5xx):**
- `500 Internal Server Error` - внутренняя ошибка сервера
- `502 Bad Gateway` - ошибка шлюза
- `503 Service Unavailable` - сервис недоступен

### 4. **Stateless (Без состояния)**
Каждый запрос должен содержать всю необходимую информацию. Сервер не хранит состояние между запросами.

### 5. **Единообразный интерфейс**
API должен быть последовательным и предсказуемым.

## Структура REST API

### Именование ресурсов

**Хорошие примеры:**
```
GET    /users              - получить всех пользователей
GET    /users/1            - получить пользователя с ID=1
POST   /users              - создать нового пользователя
PUT    /users/1            - обновить пользователя с ID=1
DELETE /users/1            - удалить пользователя с ID=1

GET    /events             - получить все события
GET    /events/5           - получить событие с ID=5
POST   /events             - создать новое событие
PATCH  /events/5           - частично обновить событие
DELETE /events/5          - удалить событие
```

**Плохие примеры:**
```
GET /getUsers              ❌ (глагол в URL)
GET /users/get/1           ❌ (глагол в URL)
POST /createUser           ❌ (глагол в URL)
GET /user_list             ❌ (непоследовательность)
```

### Вложенные ресурсы

```
GET    /users/1/events           - события пользователя с ID=1
POST   /users/1/events           - создать событие для пользователя
GET    /users/1/events/5         - конкретное событие пользователя
DELETE /users/1/events/5         - удалить событие пользователя
```

## Примеры REST API запросов

### 1. Получение ресурсов (GET)

**Получить все события:**
```http
GET /events HTTP/1.1
Host: localhost:8080
```

**Ответ:**
```json
[
  {
    "id": 1,
    "title": "Концерт",
    "description": "Описание концерта",
    "category": "Музыка"
  },
  {
    "id": 2,
    "title": "Выставка",
    "description": "Описание выставки",
    "category": "Искусство"
  }
]
```

**Получить одно событие:**
```http
GET /events/1 HTTP/1.1
Host: localhost:8080
```

**Ответ:**
```json
{
  "id": 1,
  "title": "Концерт",
  "description": "Описание концерта",
  "category": "Музыка",
  "location": {
    "city": "Москва",
    "venue": "Концертный зал"
  }
}
```

### 2. Создание ресурса (POST)

**Создать новое событие:**
```http
POST /events HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "title": "Новый концерт",
  "description": "Описание нового концерта",
  "category": "Музыка",
  "location": {
    "city": "Санкт-Петербург",
    "venue": "Филармония"
  }
}
```

**Ответ (201 Created):**
```http
HTTP/1.1 201 Created
Content-Type: application/json
Location: /events/3

{
  "id": 3,
  "title": "Новый концерт",
  "description": "Описание нового концерта",
  "category": "Музыка",
  "location": {
    "city": "Санкт-Петербург",
    "venue": "Филармония"
  }
}
```

### 3. Обновление ресурса (PUT - полное обновление)

**Полностью обновить событие:**
```http
PUT /events/1 HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "title": "Обновленный концерт",
  "description": "Новое описание",
  "category": "Музыка",
  "location": {
    "city": "Москва",
    "venue": "Новый зал"
  }
}
```

**Ответ (200 OK):**
```json
{
  "id": 1,
  "title": "Обновленный концерт",
  "description": "Новое описание",
  "category": "Музыка",
  "location": {
    "city": "Москва",
    "venue": "Новый зал"
  }
}
```

### 4. Частичное обновление (PATCH)

**Частично обновить событие:**
```http
PATCH /events/1 HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "title": "Только новое название"
}
```

**Ответ (200 OK):**
```json
{
  "id": 1,
  "title": "Только новое название",
  "description": "Старое описание",
  "category": "Музыка"
}
```

### 5. Удаление ресурса (DELETE)

**Удалить событие:**
```http
DELETE /events/1 HTTP/1.1
Host: localhost:8080
```

**Ответ (204 No Content):**
```http
HTTP/1.1 204 No Content
```

## Фильтрация, сортировка и пагинация

### Фильтрация (Query параметры)

```
GET /events?category=Музыка
GET /events?city=Москва
GET /events?category=Музыка&city=Москва
GET /events?status=active&date=2024-12-25
```

### Сортировка

```
GET /events?sort=title          - сортировка по названию
GET /events?sort=-date          - сортировка по дате (убывание)
GET /events?sort=title,date     - множественная сортировка
```

### Пагинация

```
GET /events?page=1&limit=10      - первая страница, 10 элементов
GET /events?page=2&limit=10      - вторая страница
GET /events?offset=20&limit=10  - альтернативный способ
```

**Ответ с пагинацией:**
```json
{
  "data": [
    { "id": 1, "title": "Событие 1" },
    { "id": 2, "title": "Событие 2" }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 50,
    "totalPages": 5
  }
}
```

## Обработка ошибок

### Формат ошибки

```json
{
  "error": {
    "code": "EVENT_NOT_FOUND",
    "message": "Событие с ID 999 не найдено",
    "details": {
      "id": 999,
      "resource": "events"
    }
  }
}
```

### Примеры ошибок

**404 Not Found:**
```http
HTTP/1.1 404 Not Found
Content-Type: application/json

{
  "error": {
    "code": "NOT_FOUND",
    "message": "Событие с ID 999 не найдено"
  }
}
```

**400 Bad Request:**
```http
HTTP/1.1 400 Bad Request
Content-Type: application/json

{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Неверные данные запроса",
    "details": {
      "title": "Поле 'title' обязательно для заполнения",
      "date": "Неверный формат даты"
    }
  }
}
```

**409 Conflict:**
```http
HTTP/1.1 409 Conflict
Content-Type: application/json

{
  "error": {
    "code": "DUPLICATE_EMAIL",
    "message": "Пользователь с таким email уже существует"
  }
}
```

## Версионирование API

### Через URL (рекомендуется)

```
/api/v1/events
/api/v2/events
```

### Через заголовки

```http
GET /events HTTP/1.1
Host: localhost:8080
API-Version: 2
```

## Аутентификация и авторизация

### Bearer Token (JWT)

```http
GET /events HTTP/1.1
Host: localhost:8080
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### API Key

```http
GET /events HTTP/1.1
Host: localhost:8080
X-API-Key: your-api-key-here
```

## Примеры использования в разных языках

### cURL

```bash
# GET запрос
curl http://localhost:8080/events

# GET с параметрами
curl "http://localhost:8080/events?category=Музыка&page=1"

# POST запрос
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer token123" \
  -d '{
    "title": "Концерт",
    "description": "Описание"
  }'

# PUT запрос
curl -X PUT http://localhost:8080/events/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Обновленное название"
  }'

# DELETE запрос
curl -X DELETE http://localhost:8080/events/1
```

### JavaScript (Fetch API)

```javascript
// GET - получить все события
async function getEvents() {
  const response = await fetch('http://localhost:8080/events');
  const events = await response.json();
  return events;
}

// GET - получить одно событие
async function getEvent(id) {
  const response = await fetch(`http://localhost:8080/events/${id}`);
  if (!response.ok) {
    throw new Error('Событие не найдено');
  }
  return await response.json();
}

// POST - создать событие
async function createEvent(eventData) {
  const response = await fetch('http://localhost:8080/events', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer token123'
    },
    body: JSON.stringify(eventData)
  });
  
  if (response.status === 201) {
    return await response.json();
  }
  throw new Error('Ошибка создания события');
}

// PUT - обновить событие
async function updateEvent(id, eventData) {
  const response = await fetch(`http://localhost:8080/events/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(eventData)
  });
  
  return await response.json();
}

// PATCH - частичное обновление
async function patchEvent(id, partialData) {
  const response = await fetch(`http://localhost:8080/events/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(partialData)
  });
  
  return await response.json();
}

// DELETE - удалить событие
async function deleteEvent(id) {
  const response = await fetch(`http://localhost:8080/events/${id}`, {
    method: 'DELETE'
  });
  
  if (response.status === 204) {
    return true;
  }
  throw new Error('Ошибка удаления события');
}
```

### Python (requests)

```python
import requests

BASE_URL = "http://localhost:8080"

# GET - получить все события
def get_events():
    response = requests.get(f"{BASE_URL}/events")
    return response.json()

# GET - получить одно событие
def get_event(event_id):
    response = requests.get(f"{BASE_URL}/events/{event_id}")
    if response.status_code == 404:
        raise Exception("Событие не найдено")
    return response.json()

# POST - создать событие
def create_event(event_data):
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer token123"
    }
    response = requests.post(
        f"{BASE_URL}/events",
        json=event_data,
        headers=headers
    )
    if response.status_code == 201:
        return response.json()
    raise Exception("Ошибка создания события")

# PUT - обновить событие
def update_event(event_id, event_data):
    response = requests.put(
        f"{BASE_URL}/events/{event_id}",
        json=event_data
    )
    return response.json()

# DELETE - удалить событие
def delete_event(event_id):
    response = requests.delete(f"{BASE_URL}/events/{event_id}")
    return response.status_code == 204
```

### Go (net/http)

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

const baseURL = "http://localhost:8080"

// GET - получить все события
func getEvents() ([]Event, error) {
    resp, err := http.Get(baseURL + "/events")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var events []Event
    if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
        return nil, err
    }
    return events, nil
}

// POST - создать событие
func createEvent(event Event) (*Event, error) {
    jsonData, err := json.Marshal(event)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest("POST", baseURL+"/events", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var createdEvent Event
    if err := json.NewDecoder(resp.Body).Decode(&createdEvent); err != nil {
        return nil, err
    }
    return &createdEvent, nil
}
```

## Best Practices (Лучшие практики)

### 1. Используйте правильные HTTP методы
- GET для чтения
- POST для создания
- PUT для полного обновления
- PATCH для частичного обновления
- DELETE для удаления

### 2. Используйте правильные статус коды
- 200 для успешного GET/PUT/PATCH
- 201 для успешного POST
- 204 для успешного DELETE
- 400 для ошибок валидации
- 404 для несуществующих ресурсов
- 500 для серверных ошибок

### 3. Именование ресурсов
- Используйте существительные (не глаголы)
- Используйте множественное число для коллекций
- Будьте последовательными

### 4. Формат данных
- Используйте JSON для обмена данными
- Всегда указывайте Content-Type заголовок
- Используйте единый формат дат (ISO 8601)

### 5. Обработка ошибок
- Всегда возвращайте понятные сообщения об ошибках
- Используйте стандартные HTTP статус коды
- Предоставляйте детали ошибок в структурированном формате

### 6. Безопасность
- Используйте HTTPS в продакшене
- Валидируйте все входные данные
- Используйте аутентификацию для защищенных ресурсов
- Ограничивайте частоту запросов (rate limiting)

### 7. Документация
- Документируйте все эндпоинты
- Указывайте примеры запросов и ответов
- Описывайте возможные ошибки

## Пример полного REST API для вашего проекта

### Эндпоинты для событий

```
GET    /events              - получить все события
GET    /events/:id          - получить событие по ID
POST   /events              - создать новое событие
PUT    /events/:id          - полностью обновить событие
PATCH  /events/:id          - частично обновить событие
DELETE /events/:id          - удалить событие

GET    /events?category=Музыка&page=1&limit=10  - фильтрация и пагинация
```

### Эндпоинты для пользователей

```
GET    /users               - получить всех пользователей
GET    /users/:id           - получить пользователя по ID
POST   /users               - создать нового пользователя
PUT    /users/:id           - обновить пользователя
DELETE /users/:id           - удалить пользователя

GET    /users/:id/events    - получить события пользователя
POST   /users/:id/events    - создать событие для пользователя
```

## Заключение

REST API — это стандартизированный способ создания веб-сервисов. Следуя принципам REST, вы создаете API, которое легко понимать, использовать и поддерживать.

**Ключевые моменты:**
- Ресурсы идентифицируются через URL
- HTTP методы определяют действие
- Статус коды показывают результат
- API должно быть stateless
- Используйте единообразный интерфейс
