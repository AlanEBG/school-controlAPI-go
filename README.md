# Sistema de Control Escolar - API REST

API RESTful desarrollada en Go para gestionar estudiantes, materias y calificaciones de un sistema escolar básico.

## Descripción

Este proyecto implementa una API REST completa que permite realizar operaciones CRUD (Crear, Leer, Actualizar, Eliminar) sobre tres entidades principales:
- **Estudiantes**: Gestión de información de alumnos
- **Materias**: Administración de asignaturas
- **Calificaciones**: Registro y consulta de calificaciones por estudiante y materia

## Tecnologías utilizadas

- **Go 1.21+** - Lenguaje de programación
- **Gin** - Framework web HTTP
- **GORM** - ORM para manejo de base de datos
- **SQLite** - Base de datos embebida
- **godotenv** - Gestión de variables de entorno

## Características principales

- CRUD completo para estudiantes, materias y calificaciones
- Validación de datos de entrada (email, rangos de calificaciones, etc.)
- Relaciones entre entidades con foreign keys y CASCADE
- Manejo de errores con códigos HTTP apropiados
- Respuestas en formato JSON estandarizado
- Persistencia en base de datos SQLite
- Migraciones automáticas de base de datos

## Requisitos previos

- Go 1.21 o superior instalado
- Git

## Instalación

### 1. Clonar el repositorio

```bash
git clone https://github.com/AlanEBG/school-control-api.git
cd school-control-api
```

### 2. Instalar dependencias

```bash
go mod download
```

### 3. Configurar variables de entorno (opcional)

```bash
cp .env.example .env
```

Puedes editar el archivo `.env` para cambiar el puerto o la ruta de la base de datos:

```env
PORT=8080
DB_PATH=school.db
```

### 4. Ejecutar la aplicación

```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8080`

**Salida esperada:**
```
Conexion a la base de datos exitosa
Migraciones ejecutadas correctamente
Servidor corriendo en http://localhost:8080
```

## Estructura del proyecto

```
school-control-api/
├── main.go                    # Punto de entrada de la aplicación
├── models/                    # Modelos de datos
│   ├── student.go            # Modelo de estudiante
│   ├── subject.go            # Modelo de materia
│   └── grade.go              # Modelo de calificación
├── handlers/                  # Controladores HTTP
│   ├── student.go            # Endpoints de estudiantes
│   ├── subject.go            # Endpoints de materias
│   └── grade.go              # Endpoints de calificaciones
├── database/                  # Configuración de base de datos
│   └── database.go           # Conexión y migraciones
├── router/                    # Configuración de rutas
│   └── router.go             # Definición de endpoints
├── utils/                     # Utilidades
│   ├── response.go           # Respuestas JSON estandarizadas
│   └── validator.go          # Validaciones personalizadas
├── .env.example              # Ejemplo de variables de entorno
├── .gitignore                # Archivos ignorados por Git
├── go.mod                    # Dependencias del proyecto
├── go.sum                    # Checksums de dependencias
├── postman_collection.json   # Colección de Postman
└── README.md                 # Este archivo
```

## API Endpoints

### Estudiantes

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| `POST` | `/api/students` | Crear un nuevo estudiante |
| `GET` | `/api/students` | Obtener todos los estudiantes |
| `GET` | `/api/students/:student_id` | Obtener un estudiante por ID |
| `PUT` | `/api/students/:student_id` | Actualizar un estudiante |
| `DELETE` | `/api/students/:student_id` | Eliminar un estudiante |

**Esquema de estudiante:**
```json
{
  "student_id": 1,
  "name": "string",
  "group": "string",
  "email": "string",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

### Materias

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| `POST` | `/api/subjects` | Crear una nueva materia |
| `GET` | `/api/subjects/:subject_id` | Obtener una materia por ID |
| `PUT` | `/api/subjects/:subject_id` | Actualizar una materia |
| `DELETE` | `/api/subjects/:subject_id` | Eliminar una materia |

**Esquema de materia:**
```json
{
  "subject_id": 1,
  "name": "string",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

### Calificaciones

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| `POST` | `/api/grades` | Crear una nueva calificación |
| `PUT` | `/api/grades/:grade_id` | Actualizar una calificación |
| `DELETE` | `/api/grades/:grade_id` | Eliminar una calificación |
| `GET` | `/api/grades/:grade_id/student/:student_id` | Obtener calificación específica |
| `GET` | `/api/grades/student/:student_id` | Obtener todas las calificaciones de un estudiante |

**Esquema de calificación:**
```json
{
  "grade_id": 1,
  "student_id": 1,
  "subject_id": 1,
  "grade": 95.5,
  "created_at": "datetime",
  "updated_at": "datetime",
  "student": { ... },
  "subject": { ... }
}
```

## Ejemplos de uso

### Crear un estudiante

```bash
curl -X POST http://localhost:8080/api/students \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alan Barcenas Garcia",
    "group": "A",
    "email": "alan@example.com"
  }'
```

**Respuesta exitosa (201 Created):**
```json
{
  "success": true,
  "data": {
    "student_id": 1,
    "name": "Alan Barcenas Garcia",
    "group": "A",
    "email": "alan@example.com",
    "created_at": "2025-11-23T02:30:00Z",
    "updated_at": "2025-11-23T02:30:00Z"
  }
}
```

### Crear una materia

```bash
curl -X POST http://localhost:8080/api/subjects \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Matematicas"
  }'
```

**Respuesta exitosa (201 Created):**
```json
{
  "success": true,
  "data": {
    "subject_id": 1,
    "name": "Matematicas",
    "created_at": "2025-11-23T02:31:00Z",
    "updated_at": "2025-11-23T02:31:00Z"
  }
}
```

### Crear una calificación

```bash
curl -X POST http://localhost:8080/api/grades \
  -H "Content-Type: application/json" \
  -d '{
    "student_id": 1,
    "subject_id": 1,
    "grade": 95.5
  }'
```

**Respuesta exitosa (201 Created):**
```json
{
  "success": true,
  "data": {
    "grade_id": 1,
    "student_id": 1,
    "subject_id": 1,
    "grade": 95.5,
    "created_at": "2025-11-23T02:32:00Z",
    "updated_at": "2025-11-23T02:32:00Z",
    "student": {
      "student_id": 1,
      "name": "Alan Barcenas Garcia",
      "group": "A",
      "email": "alan@example.com"
    },
    "subject": {
      "subject_id": 1,
      "name": "Matematicas"
    }
  }
}
```

### Obtener todos los estudiantes

```bash
curl http://localhost:8080/api/students
```

**Respuesta exitosa (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "student_id": 1,
      "name": "Alan Barcenas Garcia",
      "group": "A",
      "email": "alan@example.com",
      "created_at": "2025-11-23T02:30:00Z",
      "updated_at": "2025-11-23T02:30:00Z"
    }
  ]
}
```

### Obtener todas las calificaciones de un estudiante

```bash
curl http://localhost:8080/api/grades/student/1
```

**Respuesta exitosa (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "grade_id": 1,
      "student_id": 1,
      "subject_id": 1,
      "grade": 95.5,
      "subject": {
        "subject_id": 1,
        "name": "Matematicas"
      }
    }
  ]
}
```

### Actualizar un estudiante

```bash
curl -X PUT http://localhost:8080/api/students/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alan Barcenas Garcia",
    "group": "B"
  }'
```

**Nota:** La actualización es parcial, solo se modifican los campos enviados.

### Eliminar una calificación

```bash
curl -X DELETE http://localhost:8080/api/grades/1
```

**Respuesta exitosa (200 OK):**
```json
{
  "success": true,
  "data": {
    "message": "Calificacion eliminada correctamente"
  }
}
```

## Validaciones implementadas

### Estudiantes
- **Email:** Formato válido de correo electrónico
- **Nombre:** No vacío, máximo 100 caracteres
- **Grupo:** No vacío, máximo 10 caracteres
- **Email único:** No pueden existir dos estudiantes con el mismo email

### Materias
- **Nombre:** No vacío, máximo 100 caracteres

### Calificaciones
- **Calificación:** Rango válido entre 0 y 100
- **Foreign Keys:** Validación de existencia de estudiante y materia
- **Integridad referencial:** Se mantiene mediante CASCADE en la base de datos

## Códigos de estado HTTP

La API utiliza los siguientes códigos de estado:

| Código | Significado |
|--------|-------------|
| `200` | OK - Solicitud exitosa |
| `201` | Created - Recurso creado exitosamente |
| `400` | Bad Request - Datos inválidos o faltantes |
| `404` | Not Found - Recurso no encontrado |
| `500` | Internal Server Error - Error del servidor |

## Formato de respuestas

### Respuesta exitosa
```json
{
  "success": true,
  "data": { ... }
}
```

### Respuesta de error
```json
{
  "success": false,
  "error": "Mensaje descriptivo del error"
}
```

### Orden recomendado de pruebas

1. Crear varios estudiantes
2. Crear varias materias
3. Obtener todos los estudiantes
4. Crear calificaciones
5. Obtener calificaciones por estudiante
6. Actualizar datos
7. Eliminar registros


## Características adicionales implementadas

- Validación de llaves foráneas entre tablas
- Relaciones entre entidades (estudiante -> calificaciones, materia -> calificaciones)
- Persistencia en base de datos SQLite
- Validación de datos antes de insertar (formato email, rango calificaciones)
- Eliminación en cascada (al borrar estudiante se borran sus calificaciones)
- Carga de relaciones con Preload para respuestas completas
- Migraciones automáticas al iniciar la aplicación

### En caso de error de base de datos
```bash
rm school.db
go run main.go
```

### Error "ID invalido"
- Verifica que estés enviando números enteros válidos en los IDs
- Revisa la URL del endpoint

## Autor

**Alan Eduardo Barrera Gudiño**
- GitHub: [@AlanEBG](https://github.com/AlanEBG)


---

Desarrollado como proyecto final para el curso de Golang - 2025
