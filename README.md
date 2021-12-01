# REST API For Creating TODO Lists on Go

## When I was creating the project, I've improved my skills in:
- Development of Web Applications on Go, following the REST API design.
- Working with the <a href="https://github.com/gin-tonic/gin ">gin-tonic/gin</a> framework.
- Clean Architecture approach in building the application structure. Dependency injection technique.
- Working with the Postgresql database. Launching from Docker. Generation of migration files.
- Configuration of the application using <a href="https://github.com/spf13/viper ">spf13/viper</a> library. Working with environment variables.
- Working with the database using <a href="https://github.com/jmoiron/sql ">sql</a> library.
- Registration and authentication. Working with JWT. Middleware.
- Writing SQL queries.
- Graceful Shutdown

### To launch the application:

```
make build && make run
```

If the application is launched for the first time, you need to apply migrations to the database:

```
make migrate
```
