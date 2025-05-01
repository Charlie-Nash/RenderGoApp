# Go Web Application with PostgreSQL

Este es un proyecto de ejemplo en Go que se conecta a una base de datos PostgreSQL en Render y muestra una lista de vendedores en una página simple. El proyecto usa la estructura básica de Go y una plantilla HTML para renderizar la información.

## Estructura del Proyecto

La estructura del proyecto es la siguiente:

```
/go-app/
├── go.mod               # Archivo de módulo Go, que gestiona las dependencias del proyecto.
├── go.sum               # Archivo que contiene las sumas de control de las dependencias.
├── main.go              # Archivo principal con el código de la aplicación Go.
└── templates/           # Carpeta que contiene las plantillas HTML.
    └── index.html       # Plantilla HTML que se utiliza para renderizar la página de inicio.
```

### Archivos Importantes

- **go.mod**: Contiene las dependencias del proyecto y la configuración del módulo Go.
- **go.sum**: Almacena las sumas de control para las dependencias, asegurando que las versiones correctas sean utilizadas.
- **main.go**: Este es el archivo principal donde se inicializa la conexión a la base de datos PostgreSQL y se sirve la página web.
- **index.html**: Plantilla HTML que se utiliza para mostrar la información en la página web.

## Requisitos Previos

Antes de ejecutar este proyecto, asegúrate de tener instalado Go en tu máquina. Si no lo tienes, puedes instalarlo desde [https://go.dev/dl/](https://go.dev/dl/).

También necesitarás una base de datos PostgreSQL en Render (o en otro proveedor), con las credenciales de conexión correspondientes.

## Instalación

### 1. Clonar el Repositorio

Si aún no tienes el proyecto, puedes clonarlo con el siguiente comando:

```bash
git clone https://github.com/Charlie-Nash/RenderGoApp
cd go-app
```

### 2. Inicializar las Dependencias

El proyecto utiliza el paquete `github.com/lib/pq` para interactuar con la base de datos PostgreSQL. Puedes instalar las dependencias ejecutando el siguiente comando:

```bash
go mod tidy
```

### 3. Configuración de la Base de Datos

Asegúrate de tener una base de datos PostgreSQL funcionando en Render (o en otro proveedor). Debes obtener la URL de conexión proporcionada por Render.

Ejemplo de URL de conexión de PostgreSQL:

```
postgresql://<usuario>:<contraseña>@<host>:<puerto>/<nombre_base_de_datos>?sslmode=require
```

En el archivo `main.go`, reemplaza la URL de conexión por la que Render te proporcionó. Asegúrate de tener las credenciales correctas.

### 4. Ejecutar el Proyecto

Una vez configurado, puedes ejecutar el proyecto con el siguiente comando:

```bash
go run main.go
```

Esto iniciará un servidor en `localhost:8080`. Accede a esta URL en tu navegador para ver la lista de vendedores.

## Estructura de la Base de Datos

Asegúrate de tener una tabla de vendedores en tu base de datos PostgreSQL. Puedes crear la tabla utilizando el siguiente script SQL:

```sql
CREATE TABLE IF NOT EXISTS vendedores (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100),
    apellido VARCHAR(100),
    email VARCHAR(100)
);
```

## Funcionalidad

- El servidor se conecta a la base de datos PostgreSQL utilizando el paquete `github.com/lib/pq`.
- Recupera una lista de vendedores de la base de datos.
- Muestra la lista en una página HTML renderizada con la plantilla `index.html`.

## Contribuciones

Si deseas contribuir a este proyecto, por favor haz un fork y envíame un pull request con tus cambios.

## Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo `LICENSE` para más detalles.
