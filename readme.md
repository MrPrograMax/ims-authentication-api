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
  "fullname": "Linus Torwalds"
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

# Для подъема бд и запуска сервера:
```cmd
make build
make run
```
# Для добавления миграций в бд: 
```cmd
make migrate
```
## Замечание 1: Требуется зависимость make. Для установки:
Запускаете PowerShell от имени администратора
```cmd
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```
```
choco install make
```
## Замечание 2: Требуется зависимость migrate. Для установки:
```cmd
irm get.scoop.sh | iex
scoop install migrate
```
# Альтернативный способ запуска:
## Замечатние 3: Нужно самому создать и прописать файл .env:
```yaml
DB_PASSWORD:qwerty
```
Запуск:
```cmd
docker pull postgres
docker run --name=ims-database -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 --rm postgres
migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up
go run cmd/main.go
```
Для входа в postgres
```cmd
docker exec -it ims-auth-db psql -U postgres
```
