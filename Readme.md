# Broker

Broker is a complete application that investors use to control and generate signals of portifolio of assets.

## DI

Inside package of DI, create stubs using command:

```sh
wire
```

## Migration

User service
```sh
docker run --rm -v /home/marcelo/projetos/broker/backend/go-user-service/scripts/migration:/flyway/sql flyway/flyway -url=jdbc:postgresql://host.docker.internal:5432/broker_user_database -user=postgres -password=123456 migrate
```
Sts Service

```sh
docker run --rm -v /home/marcelo/projetos/broker/backend/go-user-service/scripts/migration:/flyway/sql flyway/flyway -url=jdbc:postgresql://host.docker.internal:5432/broker_user_database -user=postgres -password=123456 migrate
```
