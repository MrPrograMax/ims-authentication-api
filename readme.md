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

Для подъем бд и запуск сервера:
1. `make build`
2. `make run`

Для добавления миграций в бд: 
`make migrate`

Зачание: У вас может не быть утилиты для использования make. Тогда:
1. Запускаете PowerShell от имени администратора
2. Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
3. choco install make

Замечание. У вас может не быть утилиты migrate. Тогда:
1. irm get.scoop.sh | iex
2. scoop install migrate



