# go-shorts

Some shortcuts for go development

# Commands

## gotest

Sets up any docker-compose environment and runs the test suite for the current package / module

### Installation

```shell
go install github.com/pfouilloux/goshorts/cmd/gotest@latest
```

### Usage

```shell
gotest [-cf/-compose_file FILE] [-c/-cover] [-r/-race] [-once] <PATH>
  -cf/-compose_file FILE          = sets the docker compose file (defaults to docker-compose.yml)
  -dep/dependencies DEP1 DEP2 ... = space separated list of which services to start from docker compose, will start all services when blank or not provided 
  -c/-cover                       = show code coverage percentage (defaults to true)
  -r/-race                        = run race condition tests - requires CGO & GCC (defaults to false)
  -once                           = tears down the docker compose environment after the tests are run (defaults to false)

example: 
  > gotest -cf my_compose_file.yml -c -r -once ./...
```

##### Raw command

The -raw flag can be used to pass in custom go test command. All other testing flags will be ignored but docker & other environment flags will still
be used.

```shell
gotest [-cf/-compose_file FILE] [-raw ARGS]
  -cf/-compose_file FILE          = sets the docker compose file (defaults to docker-compose.yml)
  -dep/dependencies DEP1 DEP2 ... = space separated list of which services to start from docker compose, will start all services when blank or not provided
  -raw ARGS                       = passes the provided arguments as is to gotestsum
  -once                           = tears down the docker compose environment after the tests are run (defaults to false)
  
example: This will be the same as running "gotest -cf my_compose_file.yml -c -r ./..."
  > gotest -cf my_compose_file.yml -raw "-cover -race ./..."
```