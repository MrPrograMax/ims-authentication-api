# User Authentication Service

## 1. Авторизация

`POST` /auth/sign-in<br>
Ожидаемый формат входных данных:

```json
{
  "username": "name",
  "password": "password"
}
```

#### <span style="color:#12ff63">200 STATUS: OK
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkyMTM4NDQsImlhdCI6MTcxOTE3MDY0NCwidXNlcl9pZCI6MH0.XEuiHNUeRF1B9-0GbZoJ-M2UcPQnVuUrhJmsQ0q7gmA"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST
```json
{
  "message": "invalid request"
}
```

## 2. Регистрация

`POST` /auth/sign-up<br>
Ожидаемый формат входных данных:

```json
{
  "username": "name",
  "password": "password",
  "role_id": 1
}
```

#### <span style="color:#12ff63">200 STATUS: OK
```json
{
  "id": 1
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST
```json
{
  "message": "invalid request"
}
```

Для запуска сервера выполнить следующие действия:


1. docker pull postgres
2. docker run —name=ims-auth-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 —rm postgres
3. migrate -path ./schema 'postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

Замечание: в у вас может не быть утилиты migrate. Тогда в этом случае перед пунктом 3 выполнить шаг:

a. irm get.scoop.sh | iex
b. scoop install migrate

