# Test pg server

docker-compose.yaml

```yaml
version: '3.7'

services:
    postgres:
        image: postgres
        restart: always
        environment:
            POSTGRES_USER: docker_pg
            POSTGRES_PASSWORD: docker_pg
        ports:
            - "5432:5432"
        volumes:
            - postgres_data:/var/lib/postgresql/data

    pgadmin:
        image: dpage/pgadmin4
        restart: always
        ports:
            - "5050:80"
        environment:
            PGADMIN_DEFAULT_EMAIL: admin@example.com
            PGADMIN_DEFAULT_PASSWORD: admin

volumes:
    postgres_data:
```

Тесты проводились через postman