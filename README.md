# Tesg-Golang
## Requisitos Previos

- Docker
- Docker Compose

## Estructura del Proyecto

```
.
├── database/     # Configuración de la base de datos
├── docs/         # Documentación
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

## Ejecución

Para ejecutar el proyecto:

```bash
docker-compose up --build
```
O simplemente desde el archivo docker-compose.yml si tenemos la extension de docker y docker actualmente corriendo en nuestro ordenador
entonces le damos run all services. 

Esto iniciará:
- La aplicación Go en el puerto 8080
- MongoDB en el puerto 27017

## Detalles Técnicos

- La aplicación está construida con Go 1.24.1
- Utiliza MongoDB como base de datos
- La API está expuesta en el puerto 8080
- Los datos de MongoDB se persisten en el directorio `mongo_data`

## Detener la Aplicación

Para detener la aplicación:

```bash
docker-compose down
```