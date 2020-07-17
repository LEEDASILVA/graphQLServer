# gqlgen

you have to use go module, for example:

```console
go mod init github.com/LEEDASILVA/graphQLServer/go
```

then run and initialize the github.com/99designs/gqlgen, like:

```console
go run github.com/99designs/gqlgen init
```

**gqlgen.yml** — The gqlgen config file, knobs for controlling the generated code.

**generated.go** — The GraphQL execution runtime, the bulk of the generated code.

**models_gen.go** — Generated models required to build the graph. Often you will override these with your own models. Still very useful for input types.

**resolver.go** — This is where your application code lives. generated.go will call into this to get the data the user has requested.

**server/server.go** — This is a minimal entry point that sets up an http.Handler to the generated GraphQL server. start the server with go run **server.go** and open your browser and you should see the graphql playground, So setup is right!

# Docker Data base

For now we will use mySQL

Use docker to generate and run a Mysql image

`docker run --name mysql -e MYSQL_ROOT_PASSWORD=***** -d mysql:latest`

if we run the command `docker ps`

we can see that the docker image is running in a container

```console
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                 NAMES
63d2fc610647        mysql:latest        "docker-entrypoint.s…"   8 seconds ago       Up 6 seconds        3306/tcp, 33060/tcp   mysql
```

to add a data base lets just open the container and put sum in the db

```console
$ docker exec -it mysql bash
root@111...:/# mysql -u root -p
CREATE DATABASE hackernews;
```

# golang-migrate

```console
cd internal/pkg/db/migrations/
migrate create -ext sql -dir mysql -seq create_users_table
migrate create -ext sql -dir mysql -seq create_links_table
```

apply the migrations

```console
$ migrate -database mysql://root:****@(172.17.0.2:3306)/hackernews -path internal/pkg/db/migrations/mysql up
```

# JWT

JWT or Json Web Token is a string containing a hash that helps us verify who is using application. Every token is constructed of 3 parts like xxxxx.yyyyy.zzzzz and name of these parts are: Header, Payload and Signature.

you can see more about it [here](https://jwt.io/introduction/)
