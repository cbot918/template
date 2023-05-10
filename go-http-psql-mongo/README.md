# HTTP DB Template

### A Toy Bootstrap Api Template For golang Http and Database

<br/>

## Include

- Go-PostgreSQL

- Nodejs-MongoDB

<br/>

## Stacks

- Http: [fiber](https://github.com/gofiber/fiber)
- Database: [postgresql](https://hub.docker.com/_/postgres) / [mongodb](https://github.com/mongodb/mongo)
- Ops: [doocker-compose](https://github.com/docker/compose)

## Getting Started

### Run App

```
sudo docker-compose up --build --force-recreate
```

## Test PostgreSQL

### test api and database

```
curl localhost:5455/psql
```

### cli to connect database

```
curl localhost:5455/psql/ping
```

## Test MongoDB

```
curl localhost:5455/mongo
```

## delete app

```
sudo docker-compose down
```
