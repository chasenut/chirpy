# Chirpy

Social network platform to post *chirps*. 

## About

Chirpy manages user interaction with available content backed by secure authenticaion and authorization
features allowing creation and deletion of user created media (chirps), alongside their accounts and plans.

Users can be created, deleted, modified, authenticated, granted and revoked access to media etc. 

Written in **Go**. Based on a **RESTful API**.

## Build & run

Taking the steer of this vehicle is fairly simple.

### Tools you'll need

You will need some sort of PostgreSQL database, to which you will have to have a *connection string*.
1. [Go toolchain](https://go.dev/doc/install) `1.22+`
2. [PostgreSQL server](https://www.postgresql.org/)


### .env file

This is a special file required for the app to work. It contains 
config details and other confidential information, such as database connection string, secret and api key, etc.

Example of `.env` file:
```
DB_URL="postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable"
PLATFORM="dev"
SECRET="yLrHUfK2vtL0MiHDDTTpNN2VsqlO+v5pAn0e9FMG2xaAhpOiIAHtn0j2hKGyDxPSDizVJqLn1v5TmT99hhN03g=="
POLKA_KEY="f271c81ff7084ee5b99a5091b42d486e"
```

- **DB_URL** - connection string for database
- **SECRET** - encryption key for JWT auth tokens
- **POLKA_KEY** - just an API key

Note: `PLATFORM="dev"` enables the use of some dangerous endpoints (e.g. "POST /admin/reset"), 
it is better avoided when released *to the wild*.

### Run 

To run the app, use:
```
go run .
```

## API documentation

For API docs, see [here](./docs/API.md).

## Credit

Made as a part of [Boot.dev](https://www.boot.dev) course.
