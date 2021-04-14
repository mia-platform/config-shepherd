# Mia Config Shepherd

## Config Shepherd Command Line Tool

`config-shepherd` is a command line tool responsible for configuring a container using inside an initContainer of a k8s deployment.

The main subcommands that the tool has are:

- `joiner`: join all the same file in a list of directory and write them in an output directory


## Development Local

**build**: to build the script run: `go build ./cmd/config-shepherd/...`  
**test**: to test the script run: `go test ./... -v`

## Usage

* run `config-shepherd`  
    with:
    * `--input-dirs`: a list of directories paths containing the splitted files
    * `--output-dir`: a directory path where to put the joined final filesPath  

   or:  
    * `--config-name`: name of the config file (must be a json file) 
    * `--config-path`: path to the config file

Reminder: you can only use one of the two, and if you choose the latter your config file must follow the rule defined in `./config.schema.json`
