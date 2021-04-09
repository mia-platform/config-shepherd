# Mia Config Shepherd

## Config Shepherd Command Line Tool

`config-shepherd` is a command line tool responsible for configuring a container using inside an initContainer of a k8s deployment.

The main subcommands that the tool has are:

- `joiner`: join all the same file in a list of directory and write them in an output directory


## Development Local

**build**: to build the script run: `go build ./cmd/config-shepherd/...`  
**test**: to test the script run: `go test ./... -v`