# Texit

Texit is an api that allows for creating exit nodes in a cloud provider on demand for your tailnet.

This repo contains all code related to texit, which includes the API server, the CLI, and a Discord Bot for remote management.

Currently this project is in early stages so expect potentially breaking changes between releases.

## Whats Supported?

Texit is designed to allow for launching exit nodes on various tailnets and providers. A "tailnet" is some network that implements the Tailscale API. A "provider" is some cloud platform that is capable of running compute workloads. Currently the following are supported.

- Tailnets

  - [Tailscale](/docs/tailnets.md#tailscale)
  - [Headscale](/docs/tailnets.md#headscale)

- Providers
  - [AWS-ECS](/docs/providers.md#AWS-ECS)
  - [AWS-EC2](/docs/providers.md#AWS-EC2)

If you have a request to support a new tailnet or provider, please open an issue.

## Getting Started

For a quick setup, see the [Getting Started](docs/getting-started.md) guide for information on how to stand up the API and using the CLI.

Examples of setting up with docker can be found in the [examples](/examples/api/docker) directory.

If you want to setup the Discord Bot, see the [Discord Bot](/docs/discord-bot.md) guide.

Docker examples for the bot can also be found in the [examples](/examples/discord/docker) directory.
