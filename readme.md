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
