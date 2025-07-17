# Hotel reservation
A sample of hotel reservation or probably any reservation system in golang

# Project environment variables
```
HTTP_LISTEN_ADDRESS=:8888
JWT_SECRET=somethingsupersecret
MONGO_DB_NAME=hotel-reservation
MONGO_DB_URL=mongodb://localhost:27017
MONGO_DB_URL_TEST=mongodb://localhost:27017
```

## Project outline
- users -> book room from an hotel
- admins -> going to check reservation/bookings
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding, migration

## Resources
### Mongodb driver

Installing mongodb client
```
go get go.mongodb.org/mongo-driver/mongo
```

### go-fiber 
Documentation
```
https://gofiber.io
```

Installing go-fiber
```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a Docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```
