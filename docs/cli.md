# CLI

The Texit CLI is a command line interface that allows for interacting with the Texit API. The CLI is essentially a wrapper around this API with a few extra commands, such as `init` commands that give instructions on how to configure a tailnet or provider.

## Installation

The CLI can be installed by downloading the binary from the [release page](https://github.com/awlsring/texit/releases)

The binary tars follow the name pattern of `texit_cli_<os>_<arch>.tar.gz`. For example, the linux binary for amd64 would be `texit_cli_Linux_x86_64.tar.gz`.

After downloading the binary, you can extract it and place it in your PATH under the name `texit`.

Once on your path, you can run the following command init the CLI.

```bash
$ texit init
```

This will create a `.texit` directory in your home directory that will store your configuration and other data. A configuration file will be created at `~/.texit/config.yaml`.

## Configuration

The Texit CLI config is very simple. You must specify the address of the Texit API and the API key to use when calling the API. The following is an example configuration file.

```yaml
api:
  address: "http://localhost:7032"
  apiKey: changeme
```

## Usage

As stated before, the CLI is a wrapper around the Texit API. You can use the `--help` flag to get a list of all commands and subcommands.

```bash
$ texit --help
```
