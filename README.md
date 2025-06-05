# Proyecto Gimnasio

Este repositorio contiene un backend en Go (Gin + GORM) y un frontend en React/Vite. Permite listar y administrar actividades de un gimnasio.

## Configuración

1. Crea un archivo `.env` en `backend_xd/` con las variables:

```
DB_DSN="usuario:clave@tcp(localhost:3306)/gimnasio?charset=utf8&parseTime=true&loc=Local"
JWT_SECRET="una_clave_muy_segura"
```

2. Instala las dependencias del frontend:

```bash
cd frontend_xd
npm install
```

## Ejecución

### Backend

```bash
cd backend_xd
go run main.go
```

El servidor se inicia por defecto en `http://localhost:8000`.

### Frontend

```bash
cd frontend_xd
npm run dev
```

La aplicación se sirve en `http://localhost:5173`.

## Endpoints principales

- `GET /actividades` lista todas las actividades.
- `GET /actividad/:id` devuelve una actividad por ID.
- `PUT /actividad/:id` modifica una actividad (admin).

La página principal consume estos endpoints y permite que un administrador edite las actividades desde la UI.
