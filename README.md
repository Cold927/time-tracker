# Time Tracker

API сервиса тайм трекера

## Задача
<details><summary> Полное описание задачи можно прочесть здесь</summary>

**1. Выставить REST методы**
```markdown  
1. Получение данных пользователей:
2. Фильтрация по всем полям.
3. Пагинация. 
4. Получение трудозатрат по пользователю за период задача-сумма часов и минут с сортировкой от большей затраты к меньшей
5. Начать отсчет времени по задаче для пользователя
6. Закончить отсчет времени по задаче для пользователя
7. Удаление пользователя
8. Изменение данных пользователя
```
### Условие
   Добавление нового пользователя в формате:
```json
{
  "address": "г. Москва, ул. Ленина, д. 5, кв. 1",
  "name": "Иван",
  "patronymic": "Иванович",
  "surname": "Иванов",
  "passportNumber": "1234 567890" // серия и номер паспорта пользователя
}
```
**2. При добавлении сделать запрос в АПИ, описанного сваггером**
```yaml
openapi: 3.0.3
info:
  title: Users info
  version: 0.0.1
paths:
  /users/info:
    get:
      parameters:
        - name: passportSeries
          in: query
          required: true
          schema:
            type: integer
        - name: passportNumber
          in: query
          required: true
          schema:
            type: integer
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/defenitions/model.UserCreate'
        '400':
          description: Bad request
        '500':
          description: Internal server error
components:
  schemas:
    Users:
      required:
        - surname
        - name
        - address
      type: object
      properties:
        surname:
          type: string
          example: Иванов
        name:
          type: string
          example: Иван
        patronymic:
          type: string
          example: Иванович
        address:
          type: string
          example: г. Москва, ул. Ленина, д. 5, кв. 1
```
**3. Обогащенную информацию положить в БД postgres (структура БД должна быть создана путем миграций при старте сервиса)**

**4. Покрыть код debug- и info-логами**

**5. Вынести конфигурационные данные в .env-файл**

**6. Сгенерировать сваггер на реализованное АПИ**
</details>

## Запуск сервиса
### Если есть Docker
```
docker-compose up -d --build
```

**Ссылка на Swagger:**

http://localhost:8500/swagger/index.html


### Запуск сервиса локально

Для локального запуска, необходимо создать файл `.env.local`

```bash
# App Settings
APP_PORT=8600
APP_ENV=development
GIN_MODE=release
CORS=*

# Database
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=timetracker
DB_PORT=5432
DB_SSL_MODE=disable
DB_MAX_IDLE_CONNECTIONS=15
DB_MAX_OPEN_CONNECT=100
```

Для запуска сервиса необходимо выполнить следующие команды:
```
go mod download
```
Старт:
```
go run main.go
```
**Ссылка на Swagger:**

http://localhost:8600/swagger/index.html

## Автор

[Роман Шумилин](https://github.com/Cold927)