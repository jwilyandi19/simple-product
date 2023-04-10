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