# SQLite Команды - Справочник

## Подключение к базе данных

```bash
# Подключение к базе данных
sqlite3 test.db

# Подключение с выводом заголовков
sqlite3 -header -column test.db

# Выполнение SQL команды из командной строки
sqlite3 test.db "SELECT * FROM users;"
```

## Основные команды SQLite

### Просмотр структуры базы данных

```sql
-- Список всех таблиц
.tables

-- Список всех таблиц с типами
.table

-- Схема конкретной таблицы
.schema users
.schema events

-- Схема всех таблиц
.schema

-- Информация о таблице
PRAGMA table_info(users);
PRAGMA table_info(events);
```

### Настройки вывода

```sql
-- Включить заголовки колонок
.headers on

-- Включить режим колонок
.mode column

-- Включить режим таблицы
.mode table

-- Включить режим CSV
.mode csv

-- Включить режим JSON
.mode json

-- Включить режим списка
.mode list

-- Установить ширину колонок
.width 15 20 10

-- Включить таймер выполнения запросов
.timer on
```

### Работа с данными

#### SELECT - Выборка данных

```sql
-- Выбрать все записи из таблицы users
SELECT * FROM users;

-- Выбрать все записи из таблицы events
SELECT * FROM events;

-- Выбрать конкретные колонки
SELECT id, name, email FROM users;

-- Выбрать с условием
SELECT * FROM users WHERE id = 1;
SELECT * FROM users WHERE email = 'user@example.com';

-- Выбрать с сортировкой
SELECT * FROM users ORDER BY created_at DESC;
SELECT * FROM events ORDER BY id ASC;

-- Выбрать с лимитом
SELECT * FROM users LIMIT 10;
SELECT * FROM events LIMIT 5 OFFSET 10;

-- Подсчет записей
SELECT COUNT(*) FROM users;
SELECT COUNT(*) FROM events;

-- Поиск по тексту
SELECT * FROM events WHERE title LIKE '%концерт%';
SELECT * FROM events WHERE description LIKE '%музыка%';
```

#### INSERT - Вставка данных

```sql
-- Вставка пользователя
INSERT INTO users (name, email, created_at, updated_at) 
VALUES ('Иван Иванов', 'ivan@example.com', datetime('now'), datetime('now'));

-- Вставка события (пример)
INSERT INTO events (
    title, 
    description, 
    short_description,
    location_venue,
    location_address,
    location_city,
    location_country,
    date_start_date,
    date_end_date,
    date_start_time,
    date_end_time,
    date_timezone,
    organizer_name,
    organizer_email,
    organizer_phone,
    category,
    tags,
    image_url,
    image_alt,
    visibility,
    metadata_created_at,
    metadata_updated_at
) VALUES (
    'Концерт',
    'Описание концерта',
    'Краткое описание',
    'Концертный зал',
    'ул. Примерная, 1',
    'Москва',
    'Россия',
    '2024-12-25',
    '2024-12-25',
    '19:00',
    '22:00',
    'Europe/Moscow',
    'Организатор',
    'org@example.com',
    '+79991234567',
    'Музыка',
    '["концерт", "музыка"]',
    'https://example.com/image.jpg',
    'Концерт',
    'public',
    datetime('now'),
    datetime('now')
);
```

#### UPDATE - Обновление данных

```sql
-- Обновить пользователя
UPDATE users 
SET name = 'Новое имя', updated_at = datetime('now') 
WHERE id = 1;

-- Обновить событие
UPDATE events 
SET title = 'Новое название', updated_at = datetime('now') 
WHERE id = 1;

-- Обновить несколько полей
UPDATE users 
SET name = 'Имя', email = 'newemail@example.com', updated_at = datetime('now') 
WHERE id = 1;
```

#### DELETE - Удаление данных

```sql
-- Удалить пользователя
DELETE FROM users WHERE id = 1;

-- Удалить событие
DELETE FROM events WHERE id = 1;

-- Удалить все записи (осторожно!)
DELETE FROM users;
DELETE FROM events;
```

### Работа с таблицами

#### Создание таблиц

```sql
-- Создать таблицу users (пример структуры)
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    created_at DATETIME,
    updated_at DATETIME
);

-- Создать таблицу events (пример структуры)
CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    short_description TEXT,
    location_venue TEXT,
    location_address TEXT,
    location_city TEXT,
    location_country TEXT,
    date_start_date TEXT,
    date_end_date TEXT,
    date_start_time TEXT,
    date_end_time TEXT,
    date_timezone TEXT,
    organizer_name TEXT,
    organizer_email TEXT,
    organizer_phone TEXT,
    category TEXT,
    tags TEXT,
    image_url TEXT,
    image_alt TEXT,
    visibility TEXT,
    metadata_created_at TEXT,
    metadata_updated_at TEXT
);
```

#### Изменение таблиц

```sql
-- Добавить колонку
ALTER TABLE users ADD COLUMN phone TEXT;

-- Переименовать таблицу (SQLite не поддерживает напрямую)
-- Нужно создать новую таблицу, скопировать данные, удалить старую

-- Удалить таблицу
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS events;
```

### Индексы

```sql
-- Создать индекс
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_events_category ON events(category);
CREATE INDEX idx_events_title ON events(title);

-- Показать все индексы
.indices users
.indices events

-- Удалить индекс
DROP INDEX idx_users_email;
```

### Транзакции

```sql
-- Начать транзакцию
BEGIN TRANSACTION;

-- Коммит транзакции
COMMIT;

-- Откат транзакции
ROLLBACK;
```

### Экспорт и импорт данных

```sql
-- Экспорт в CSV
.mode csv
.headers on
.output users.csv
SELECT * FROM users;
.output stdout

-- Экспорт в SQL
.output backup.sql
.dump users
.dump events
.output stdout

-- Импорт из SQL файла
.read backup.sql

-- Импорт из CSV
.mode csv
.import users.csv users
```

### Резервное копирование

```bash
# Создать резервную копию базы данных
sqlite3 test.db ".backup backup.db"

# Восстановить из резервной копии
sqlite3 backup.db ".backup test.db"
```

### Полезные команды

```sql
-- Показать версию SQLite
SELECT sqlite_version();

-- Показать информацию о базе данных
.databases

-- Показать все настройки
.show

-- Выход из SQLite
.quit
.exit

-- Очистить экран
.shell clear

-- Выполнить SQL из файла
.read script.sql

-- Сохранить результат запроса в файл
.output results.txt
SELECT * FROM users;
.output stdout
```

### Специфичные запросы для проекта

```sql
-- Найти все события определенной категории
SELECT * FROM events WHERE category = 'Музыка';

-- Найти события по городу
SELECT * FROM events WHERE location_city = 'Москва';

-- Найти события в определенном диапазоне дат
SELECT * FROM events 
WHERE date_start_date >= '2024-12-01' 
AND date_start_date <= '2024-12-31';

-- Найти пользователей, созданных за последний месяц
SELECT * FROM users 
WHERE created_at >= datetime('now', '-1 month');

-- Подсчет событий по категориям
SELECT category, COUNT(*) as count 
FROM events 
GROUP BY category;

-- Подсчет событий по городам
SELECT location_city, COUNT(*) as count 
FROM events 
GROUP BY location_city;
```

### Прагмы (настройки SQLite)

```sql
-- Включить внешние ключи
PRAGMA foreign_keys = ON;

-- Показать информацию о внешних ключах
PRAGMA foreign_key_list(events);

-- Показать информацию о таблице
PRAGMA table_info(users);

-- Показать список всех таблиц
PRAGMA table_list;

-- Оптимизация базы данных
PRAGMA optimize;

-- Проверка целостности базы данных
PRAGMA integrity_check;

-- Быстрая проверка
PRAGMA quick_check;

-- Показать размер базы данных
PRAGMA page_count;
PRAGMA page_size;
```

### Работа с JSON (для поля tags)

```sql
-- Извлечь JSON данные (SQLite 3.38+)
SELECT id, title, json_extract(tags, '$[0]') as first_tag 
FROM events;

-- Проверить, является ли значение валидным JSON
SELECT id, title, json_valid(tags) as is_valid_json 
FROM events;

-- Подсчет элементов в JSON массиве
SELECT id, title, json_array_length(tags) as tag_count 
FROM events;
```
