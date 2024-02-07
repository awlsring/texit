# Getting Started

## Pre-requisites

Before following this guide, you will need to have installed the CLI. You can download the latest release from the releases page.

Commands in this guide will assume you have placed the binary in your PATH under the "texit" name.

## Setup

Texit has a concept of "Tailnets" and "Providers". Tailnets are networks that exit nodes can be placed in. Providers are cloud platforms that the exit nodes can be launched on. You must have at least 1 tailnet and 1 provider configured to launch the api.

You can read more about [tailnets](/docs/tailnets.md) and [providers](/docs/providers.md) and how to configure them in their respective sections.

### Configuration

The API server is configured using a YAML file. The following is an example configuration file:

```yaml
server:
  address: :7032
  apiKey: asecurekeyusedtoaccesstexit
tailnets:
  - apiKey: "tskey-api-keep-me-secret"
    tailnet: "user.service"
    type: "tailscale"
    user: "user@service"
providers:
  - type: aws-ecs
    accessKey: "abcdthisnotakey"
    secretKey: "hijkdontletpeopleknowme"
    name: "my-provider"
database:
  engine: sqlite
```

This files specifies the following:

- The server will be launched on port 7032 and will require the api key `asecurekeyusedtoaccesstexit` to be passed in `X-Api-Key` header.
- The tailnet `user.service` will be a tailnet that exit nodes can be placed in.
- The provider `my-provider` will be an AWS-ECS provider that will be used to launch the exit nodes.
- SQLite will be used as the database engine.

Before adding your Tailnets or Providers to this configuration, you must first configure them to grant Texit access to create resources.

### Initing Providers and Tailnets

Before you can target a Tailnet or a Provider, you must first initialize them with the need configuration so that Texit can access them. The Texit CLI allows for `init` commands that will give a walkthrough of how to configure the Tailnet or Provider.

The following command will give you steps of how to configure the AWS-ECS provider:

```bash
$ texit provider init -t aws-ecs
```

The following command will give you steps of how to configure the Tailscale tailnet:

```bash
$ texit tailnet init -t tailscale
```

### Running the API

You can run the API via it go binary or via docker. The following is an example of running the API server via docker:

```bash
$ docker run -v /path/to/config.yaml:/config.yaml -v /path/to/data:/var/lib/texit -p 7032:7032 ghcr.io/awlsring/texit:latest
```

This command will use your config file, store the SQLite database in `/path/to/data`, and bind the port to 7032 on your host.

Once this is up, your texit server will be running and ready to accept requests. You validate that the server is running by using the `check-health` command with the CLI:

```bash
$ texit health
```

### Using the API

Texit exposes various API methods that let you create, start, stop, and delete exit nodes. This is done with a REST api can be accessed via curl, or http clients or the Texit CLI.

The Texit CLI is written to have methods that map to the "restful" endpoints of the API. All API operations are supported as commands.

You can view more information on available CLI commands by running `texit --help`.

#### Creating an Exit Node

The following command will create an exit node in your cluster:

```bash
$ texit node provision --provider <YOUR_PROVIDER> --location <A LOCATION> --tailnet <YOUR_TAILNET>
```

Running this will launch a workflow that will create your node on your provider and add it to your tailnet. The CLI by default will poll this workflow execution until completion.

Once completed, you should be able to see the exit node listed in your devices on your tailnet

#### Stopping and Starting an Exit Node

Since you don't always need your exit node running, you can stop it with the following command:

```bash
$ texit node stop --id <YOUR_NODE_ID>
```

This will stop the exit node. When you want to use it again, you can start it with the following command:

```bash
$ texit node start --id <YOUR_NODE_ID>
```

#### Deleting an Exit Node

If you no longer need an exit node, you can delete it with the following command:

```bash
$ texit node delete --id <YOUR_NODE_ID>
```

This will delete the exit node from your provider and remove it from your tailnet.
