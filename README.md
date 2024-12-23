# Gin-HTMX boilerplate

This is a boilerplate repository for developing web apps with Go, Gin, Postgres, HTMX, and Tailwindcss.

## Project layout
The structure of this project is based on the (standard golang project layout)[https://github.com/golang-standards/project-layout].

## Development

### Setup

- Install Go
- Install Docker
- Install (sqlc)[https://docs.sqlc.dev/en/stable/overview/install.html#go-install]
- Install node

```bash
# Install npm dependencies
npm install

# Run the postgres container
docker compose up -d

# Create the .env file according to the template
cp .env.template .env

# Run the database migrations
dbmate up
```

There are a couple of ways to start up the application: Using Gin debug or release modes. The main difference as of now, is how the application includes the javascript and css in [base.html](/internal/templates/components/base.html). Debug mode uses vite which does the building in memory and serves the files from a proxy server at localhost:8081, wheras release mode has vite build the files to `/internal/static/build`, which are served from there.

**Starting up with debug mode**
```bash
npm run dev

export GIN_MODE=debug   # This is the default behaviour
go run cmd/main.go
```

**Starting up with release mode**
```bash
# Transpile the ts and build css once
npm run build:prod
# Or alternatively, watch for changes
npm run build:dev 

export GIN_MODE=release
go run cmd/main.go
```

### Migrations

The project uses dbmate to manage migrations. See the [dbmate docs](https://github.com/amacneil/dbmate?tab=readme-ov-file#usage). Note that dbmate relies on environment variables for configuration, which it can also read from a `.env` file.


```bash
# dbmate is a development dependency in this project and should already be installed.
# There are alternative ways of installing dbmate. For example:
npm install -g dbmate

dbmate --help   # print usage help
dbmate new      # generate a new migration file
dbmate up       # create the database (if it does not already exist) and run any pending migrations
dbmate create   # create the database
dbmate drop     # drop the database
dbmate migrate  # run any pending migrations
dbmate rollback # roll back the most recent migration
dbmate down     # alias for rollback
dbmate status   # show the status of all migrations (supports --exit-code and --quiet)
dbmate dump     # write the database schema.sql file
dbmate load     # load schema.sql file to the database
dbmate wait     # wait for the database server to become available
```

### sqlc

Add new queries inside the `internal/db/queries` directory. Then, use `sqlc` to generate the go code to use the queries in go.

```bash
sqlc compile    # Statically check SQL for syntax and type errors
sqlc generate   # Generate source code from SQL
```
