# Texit

Texit is an api that allows for creating exit nodes in a cloud provider on demand for your tailnet.

This repo contains code related to texit, which includes the API server, the CLI, and a Discord Bot for remote management.

Currently this project is in early stages so expect potentially breaking changes between releases.

## Whats Supported?

Texit is designed to allow for launching exit nodes on various tailnets and providers. A "tailnet" is some network that implements the Tailscale API. A "provider" is some cloud platform that is capable of running compute workloads. Currently the following are supported.

- Tailnets

  - [Tailscale](/docs/tailnets.md#tailscale)
  - [Headscale](/docs/tailnets.md#headscale)

- Providers
  - [AWS-ECS](/docs/providers.md#AWS-ECS)
  - [AWS-EC2](/docs/providers.md#AWS-EC2)
  - [Linode/Akamai](/docs/providers.md#linodeakamai)
  - [Hetzner](/docs/providers.md#hetzner)

If you have a request to support a new tailnet or provider, please open an issue.

## Getting Started

For a quick setup to run a local Texit instance, see the [getting started](docs/getting-started.md) guide for information on how to stand up the API and how to interact with it via the CLI.

An optional Discord bot is also available for remote management. See the [Discord Bot](docs/discord-bot.md) guide for more information.

Texit and the Discord bot can be deployed as a serverless application to AWS. See the [serverless](docs/serverless/serverless.md) guide for more information.
