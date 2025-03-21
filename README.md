# Tabla de contenidos
1. [Solucion](#solución)
2. [Desafíos](#desafíos)
3. [Requisitos](#requisitos)
4. [Testeo](#testeo)
5. [Instrucciones de ejecución](#instrucciones-de-ejecución)
6. [Base de Datos](#y-la-base-de-datos)

## Solución

La solución planteada está desarrollada en **go**, utilizando la librería *gorilla-mux* para el manejo de rutas de la API y utilizando **PostgreSQL** para la base de datos.

## Desafíos

El principal desafío del proyecto fue crear una API pequeña sin dejar de lado la posibilidad de que deba crecer en em futuro y, en consecuencia, desarrollar codigo escalable y extensible. 

## Requisitos

El proyecto se puede correr enteramente desde docker. Por lo tanto, el único requisito para levantar el proyecto es tener instalado docker.

## Testeo

Para los tests se utilizó la biblioteca nativa de testeo de **go**, [https://pkg.go.dev/testing](*testing*).

Para correr los tests, se puede ejecutar el siguiente comando desde la raíz del proyecto:

```
go test -v ./...
```

## Instrucciones de ejecución

Para levantar el proyecto, solo se debe correr el siguiente comando en el directorio principal:
```
docker compose up -d --build
```

### Y la base de datos?

La base de datos no es necesario levantarla aparte, ya que el docker-compose ya se encarga de eso. Además, el repositorio cuenta con un archivo `init.sql` que se ejecuta al iniciar el servicio y crea las tablas necesarias para correr (solo es una, la tabla `courses`).
