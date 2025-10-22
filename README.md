# Sirius Backend (Go + Echo + Postgres)

README Sistema de mensajería multicanal (Discord/Slack) escrito en Go con Echo + GORM + PostgreSQL. Incluye seguimiento de estado por servicio, JWT y desarrollo local con Docker.

# Resumen

Este servicio centraliza el envío de mensajes a múltiples proveedores, Discord y Slack. A su vez registra los estados del envio de mensaje por proveedor con un estado claro (success | failed). Expone una API REST con autenticación por JWT.

# Flujo Tipico:

El cliente se autentica y obtiene un token mediante JWT.

El cliente crea un mensaje indicando los servicios destino.

El backend envía por servicio y crea un registro en MessageDestination por servicio con su status.

Endpoints de admin exponen conteos diarios/totales por usuario.

Características

✅ Mensajería multicanal: Discord, Slack (Webhook/Bot API)

✅ Seguimiento por servicio con tabla message_destinations

✅ JWT con soporte de roles

✅ Métricas de administración: totales y de hoy por usuario

✅ Swagger/OpenAPI para documentación e interacción

✅ Docker Compose para DB y herramientas (PostgreSQL, pgAdmin)

# Stack Tecnologico

  Lenguaje: Go ≥ 1.22.

  Web Framework: Echo.

  GORM: GORM (PostgreSQL).

  Base de datos: PostgreSQL.

  Contenedores: Docker / Docker Compose.

# Puesta en Marcha
  Requisitos
    Go ≥ 1.22.
    Docker Desktop (con WSL2/Virtual Machine Platform habilitado en Windows).
  Git.

# Variables de entorno

Copiá .env.example a .env.dev.

Completar al menos:

ENVIRONMENT=dev.

PORT=8181.

JWT_SECRET=<tu_clave_segura>.

DISCORD_WH=<webhook>.

SLACK_WH=<webhook>.

DB (si usás Docker Compose): POSTGRES_HOST=db, POSTGRES_USER=postgres, POSTGRES_PASSWORD=postgres, POSTGRES_DB=appdb_prod, POSTGRES_PORT=5432, POSTGRES_SSL=disable.

Local sin Docker: POSTGRES_HOST=localhost - POSTGREES_PORT=5433.

Evitá espacios alrededor del = (ej.: SLACK_WH= y no SLACK_WH =). No pongas comillas a números.

# Arranque Rápido (Docker)

1) Prepará el entorno

  Copiá .env.example a .env.dev y completá valores.

  Si usás Docker, dejá POSTGRES_HOST=db, ENVIRONMENT=dev y PORT=8181.

2) Levantar todo
  docker compose up --build -d.

  La API corre dentro del contenedor en :8181.
  El puerto expuesto en tu máquina es 1000 → accedés a la API por http://localhost:1000.
  
3) Cambiar puertos

Para cambiar el puerto externo: editá el lado izquierdo en "1000:8181" (por ej., "8080:8181").

Para cambiar el puerto interno (donde escucha la app): cambiá PORT en el env y el lado derecho del mapeo (por ej., "1000:9090").

4) Comandos útiles
  docker compose down # bajar todo.
  docker compose up -d --build.

# Correr local (sin Docker)
  Go mod tidy.
  POSTGRES_HOST=localhost.
  POSTGREES_PORT=5433.
  http://localhost:8181.
  

# Tests dentro del contenedor
  docker compose exec api go test ./... -v

# Flujo de Autenticación

Register -> Login con username o email -> JWT.

Enviar Authorization: Bearer <token> a endpoints protegidos.

Roles: por defecto user; podés agregar admin para endpoints de métricas.

## Endpoints

**Base URL:** `http://localhost:1000`  
> Usa `{{token}}` (Bearer) para los endpoints protegidos.

### Login
- **URL:** `/auth/login`
- **Method:** `POST`
- **Body:**
```json
{
  "username": "siriusadmin",
  "password": "admin123!"
}
```
respuesta: "<jwt_token_string>".

### Registrar Usuario con Token

- **URL:** `url/loged/reguser`
- **Method:** `POST`
- **Auth**: Bearer {{token}}.
- **Body:**
```json
{
  "username": "admin_test1",
  "email": "admin_test1@example.com",
  "password": "password123"
}
```
- **Repuesta:** "Usuario Eliminado Correctamente" 

### CheckStatus con token
- **URL:** `url/health`
- **Method:** `POST`
- **Auth**: Bearer {{token}}.
- **Repuesta:** "OK"

### Metricas
- **URL:** `url/admin/metrics`
- **Method:** `GET`
- **Auth**: Bearer {{token}}.
- **Repuesta:** "Lista por usuario de los Mensajes enviados y restates. Cantidad de mensajes enviados"

### Eliminar Usuario con Token

- **URL:** `url/loged/deluser`
- **Method:** `DELETE`
- **Auth**: Bearer {{token}}.
- **Body:**
```json
{
  "username": "admin_test3",
  "email": "admin_test3@example.com"
}
```
- **Repuesta:** "Usuario Eliminado Correctamente"

### Actualizar Role de un Usuario

- **URL:** `url/loged/upduser`
- **Method:** `PATCH`
- **Auth**: Bearer {{token}}.
- **Body:**
```json
{
  "username": "admin_test2",
  "email": "admin_test2@example.com",
  "role": "admin"
}
```
- **Repuesta:** "Usuario Actualizado Correctamente"

### Enviar Mensaje

- **URL:** `url/message/send`
- **Method:** `POST`
- **Auth**: Bearer {{token}}.
- **Body:**
```json
{
  "content": "Prueba Discord y Slack",
  "services": [
    { "app": "discord" },
    { "app": "slack" }
  ]
}
```
```json
{
  "content": "Prueba Discord",
  "services": [
    { "app": "discord" },
  ]
}
```
- **Repuesta:** "Mensaje Enviado Correctamente"

### Obtener todos los mensajes

- **URL:** `url/admin/list`
- **Method:** `GET`
- **Auth**: Bearer {{token}}.
- **Repuesta:** "Una lista de todos los mensajes enviados con usuario y contenido"

### Obtener Messages by Status and Services (hoy)

- **URL:** `url/message/today/?status=failed&service=slack,discord`
- **Method:** `GET`
- **Auth**: Bearer {{token}}.

**Notas**: status ∈ {failed,success} — service ∈ {slack,discord} o ambos separados por coma
Ambos casos funcionan con status=failed o status=success

- **Repuesta:** "Una lista de todos los mensajes enviados del usuario en el dia"
  
### Get Messages by Status/Service Between Dates

- **URL:** `url/message/date/?status=failed&service=slack,discord&between=2025-10-01,2025-10-20`
- **Method:** `GET`
- **Auth**: Bearer {{token}}.

**Notas**: 
between=YYYY-MM-DD,YYYY-MM-DD (primera fecha desde, segunda hasta, inclusives).
service=slack,discord = Trae los estados de los mensajes de ambos servicios.
service=slack = Trae los estados de los mensaje de un servicio.
- **Repuesta:** "Una lista de todos los mensajes enviados del usuario en las fechas dadas"
