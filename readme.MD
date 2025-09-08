# go-grpc-crud

Microservicio en **Go** que implementa un **CRUD simple** con **PostgreSQL** utilizando **gRPC** como protocolo de comunicación.  
Incluye configuración parametrizada, soporte para contenerización con Docker, migraciones de base de datos con Flyway, y está diseñado pensando en despliegues en Kubernetes.

---

## 🚀 Tecnologías utilizadas

- **Go 1.22+**
- **gRPC**
- **Protocol Buffers (protoc)**
- **GORM** (ORM para Go)
- **PostgreSQL**
- **Viper** (manejo de configuración)
- **Flyway** (migraciones de base de datos)
- **Docker / Docker Compose**
- **Kubernetes (opcional, para despliegue futuro)**
- **Resiliencia** con patrones como timeouts, retries y circuit breakers.

---

## 📂 Estructura del proyecto

```bash
go-grpc-crud/
├── cmd/                # entrypoints (main.go)
├── config/             # configuración YAML (config.yaml, config.example.yaml)
├── migrations/                 # migraciones (Flyway)
├── internal/           
│   ├── config/         # carga de configuración con Viper
│   ├── db/             # conexión a la base de datos
│   ├── model/          # modelos GORM
│   ├── repository/     # capa de acceso a datos
│   ├── service/        # lógica de negocio + implementación gRPC
├── proto/              # archivos .proto y .pb.go generados
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## ⚙️ Configuración

El proyecto utiliza Viper para leer la configuración desde un archivo config.yaml o variables de entorno.
Ejemplo de config/config.yaml:

```yaml
database:
  host: localhost
  port: 5432
  user: postgres
  password: postgres
  name: mydb
  sslmode: disable
```
No se debe versionar este archivo con credenciales reales.
En su lugar, subir config/config.example.yaml con valores dummy.

## 🗄️ Migraciones de base de datos

Se utiliza Flyway para manejar las migraciones de PostgreSQL.

Estructura de migraciones (db/migrations):

```bash
migrations/
├── V1__init.sql
├── V2__add_users_table.sql
```

Ejemplo de comando con Docker Compose:
```bash
docker-compose run --rm flyway migrate
```

## 🛠️ Generación de código gRPC
#### Nota: Para este caso se debe de tener protoc como CLI, en el caso no se tenga, instalarlo desde "Chocolate" con el siguiente comando
```bash
choco install protoc
```

Si modificas los archivos .proto, puedes regenerar los stubs ejecutando:

```bash
protoc --go_out=. --go-grpc_out=. proto/user.proto
```
En el repo se versionan tanto los .proto como los .pb.go generados para evitar depender de que cada dev tenga instalado protoc.

## 🐳 Contenerización
### Dockerfile

El proyecto cuenta con un Dockerfile multistage para construir la aplicación en Go y ejecutarla en Alpine.

### Docker Compose

Archivo docker-compose.yml con servicios:

- db → PostgreSQL
- app → microservicio Go
- flyway → contenedor de migraciones

Ejemplo:
```bash
docker-compose up --build
```

## ▶️ Ejecución local

Instalar dependencias:
```bash
go mod tidy
```

Cargar configuración:

Crear archivo config/config.yaml basado en config/config.yaml.

Levantar PostgreSQL con Docker Compose:
```bash
docker-compose up -d db
```

Aplicar migraciones:
```bash
docker-compose run --rm flyway migrate
```

Ejecutar el microservicio:
```bash
go run cmd/main.go
```

El servidor gRPC estará disponible en localhost:50051.

## 🧪 Pruebas con Postman

Puedes probar el servicio con: 
- Enter URL: localhost:50051
- Services definition: Cargar el archivo /proto/user.proto