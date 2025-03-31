# Tesg-Golang

API REST para la gestión de eventos desarrollada en Go. Esta aplicación permite crear, leer, actualizar y eliminar eventos, además de proporcionar funcionalidades adicionales para la gestión de eventos.

## Descripción del Proyecto

Este proyecto es una API REST que proporciona endpoints para:
- Crear eventos
- Obtener lista de eventos
- Obtener un evento específico
- Actualizar eventos
- Eliminar eventos
- Verificar estado de eventos

La aplicación está construida con:
- Go 1.24.1
- Gin (Framework web)
- MongoDB (Base de datos)
- Swagger (Documentación API)

## Requisitos

- Docker y Docker Compose
- Go 1.24.1 o superior (para desarrollo local)

## Estructura del Proyecto

```
.
├── database/     # Configuración de la base de datos
├── docs/         # Documentación Swagger
├── handlers/     # Manejadores de la API
├── models/       # Modelos de datos
├── utils/        # Utilidades
├── .env          # Variables de entorno
├── docker-compose.yml
└── Dockerfile
```

## Configuración

1. Crea un archivo `.env` en la raíz del proyecto con las siguientes variables:

```env
MONGO_INITDB_ROOT_USERNAME=tu_usuario
MONGO_INITDB_ROOT_PASSWORD=tu_contraseña
MONGO_URI=mongodb://tu_usuario:tu_contraseña@mongo:27017
```

## Construcción y Ejecución

Para construir y ejecutar el proyecto:

```bash
docker-compose up --build
```
O simplemente desde el archivo docker-compose.yml si tenemos la extension de docker y docker actualmente corriendo en nuestro ordenador
entonces le damos run all services. 

Esto iniciará:
- La aplicación Go en el puerto 8080
- MongoDB en el puerto 27017

## Acceso a la API

La API está disponible en `http://localhost:8080` con los siguientes endpoints:

### Endpoints Disponibles

- `POST /events` - Crear un nuevo evento
- `GET /events` - Obtener todos los eventos
- `GET /events/:id` - Obtener un evento específico
- `PUT /events/:id` - Actualizar un evento
- `DELETE /events/:id` - Eliminar un evento
- `GET /events/check` - Verificar estado de eventos

## Documentación Swagger

La documentación interactiva de la API está disponible en:
```
http://localhost:8080/swagger/index.html
```

## Ejemplos de Uso

### Crear un Evento
```bash
curl -X POST http://localhost:8080/events \
-H "Content-Type: application/json" \
-d '{
    "title": "Conferencia de Go",
    "description": "Conferencia sobre desarrollo en Go",
    "date": "2024-04-01T10:00:00Z",
    "location": "Sala Principal"
}'
```

### Obtener Todos los Eventos
```bash
curl http://localhost:8080/events
```

### Obtener un Evento Específico
```bash
curl http://localhost:8080/events/ID_DEL_EVENTO
```

### Actualizar un Evento
```bash
curl -X PUT http://localhost:8080/events/ID_DEL_EVENTO \
-H "Content-Type: application/json" \
-d '{
    "title": "Conferencia de Go Actualizada",
    "description": "Nueva descripción",
    "date": "2024-04-02T10:00:00Z",
    "location": "Nueva Sala"
}'
```

### Eliminar un Evento
```bash
curl -X DELETE http://localhost:8080/events/ID_DEL_EVENTO
```

### Verificar Estado de Eventos
```bash
curl http://localhost:8080/events/check
```

## Detener la Aplicación

Para detener la aplicación:

```bash
docker-compose down
```