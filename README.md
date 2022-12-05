## Smart Loader service is API which helps you to download too many files from the Internet

###
- Postgres
- Liquibase
- GO 1.19.3
- Gorilla MUX
- Nats

For application need EnvFile by Borys Pierov plugin and .env file which contains:
```dotenv
HOST=host.docker.internal

HTTP_HOST=#{HOST}
HTTP_PORT=[your application port here]

POSTGRES_VERSION=15
POSTGRES_HOST=host.docker.internal
POSTGRES_PORT=[your postgres port here]
POSTGRES_DB=smart_loader
POSTGRES_SCHEMA=smart_loader
POSTGRES_URL=jdbc:postgresql://${HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?currentSchema=${POSTGRES_SCHEMA}
POSTGRES_USER=[your postgres user here]
POSTGRES_PASSWORD=[your postgres password here]

LIQUIBASE_VERSION=4.17

NATS_VERSION=2.9.8
NATS_HTTP_PORT=8222
NATS_PORT=4222
NATS_URL=nats://${HOST}:${NATS_PORT}
```

For successfully running liquibase need to append in db/liquibase.properties:
```dotenv
username: [your postgres user here]
password: [your postgres password here]
```

Command for building application
```dotenv
- make build
```

Command for running docker containers
```dotenv
- make docker
```