# Unity Cloud Build CLI

![Unity Cloud Build](docs/Unity_Technologies_logo.svg.png)

Command line tool for interacting with Unity Cloud Build. 

Aimed at making your CI/CD life just a little easier.

## Build

```bash
$ go get ; go build -o ucb
```

## Usage

### Configure Project Settings

You can pass these variables in as flags but that gets to be painful.

```bash
$ export UCB_API_KEY=xxx
$ export UCB_PROJECT_ID=yyy
$ export UCB_ORG_ID=zzz
```

### Get Help

```bash
$ ./ucb  help
ucb is a CLI for Unity Cloud build

Usage:
  ucb [flags]
  ucb [command]

Available Commands:
  create      Create a new object
  delete      Delete a new object
  download    Download files
  get         Get an object
  help        Help about any command
  list        List objects
  update      Update an object
  watch       Wait and watch

Flags:
  -h, --help             help for ucb
  -k, --key string       Unity Cloud Build API Key (required) (default "f68d89be5089247de212e0343db5d2e8")
  -o, --org string       Unity Cloud Build Org Id (required) (default "tapthereinc")
  -p, --project string   Unity Cloud Project Id (required) (default "42a56f22-c459-4403-b381-6599ff333803")

Use "ucb [command] --help" for more information about a command.
```

### List Build Targets

```bash
$ ucb list targets
```

### Get Single Build

```bash
$ ucb get build -t <target> <build_num>
```

### Download Build Artifacts

```bash
$ ./ucb download build -t <target> <build_num>
```

## Status

This is a work in progress. I have been adding features as I need them without aiming to be truly comprehensive.