# OBGS

This is the backend part of OBGS. The main technologies used here are:

- [ent](https://entgo.io/) for entity management
- [gqlgen](https://gqlgen.com/) for code generation
- [viper](https://github.com/spf13/viper) and [cobra](https://github.com/spf13/cobra) for CLI commands
- [atlas](https://atlasgo.io/) for managing Postgres migrations
- [taskfile](https://taskfile.dev/) for running development scripts

## Development

To start working on this you will need:

- [go](https://go.dev/)
- `go install ariga.io/atlas/cmd/atlas@latest`
- `go install github.com/go-task/task/v3/cmd/task@latest`
- [docker](https://www.docker.com/)

Once that is set up you can use the docker-compose file to run Postgres and Minio locally:

```shell
docker-compose up -d
```

Run the Postgres migrations:

```shell
task atlas_apply
```

Optionally, you can seed the DB by running `task seed`. You will need to have `psql` installed locally if you want to do this.

After all that, simply run `go run main.go server` to start listening for connections. You are good to go!
