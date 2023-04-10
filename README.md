# Simple Product

## How to Run

Run Database. This will create docker image for mysql and create new database (simple-product)
```
make run-db
```

Run Redis. This will create redis image
```
make run-redis
```

Migrate all tables to be created in DB
```
make migrate-up
```

Run program
```
make start
```
### Why Clean Architecture?
Because with clean architecture, you don't need to depend on the framework, and thus migration of framework will be much easier. Also for database layer and application layer will be separated, so database migration will be easier.

### How to scale app?
Either add more instances, or implementing Replica Database, thus request will not be centered in one database.
