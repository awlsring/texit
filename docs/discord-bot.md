# Discord Bot

Texit has a discord bot that can be used to call your Texit API. The bot is based off the same core codebase as the CLI and offers the same functionality. The bot has some additional feature such as auto completion of commands and the ability to run commands in a channel or a private message.

## Prerequisites

To install this bot, you must be able to do the following.

- Create a discord bot on the [discord developer portal](https://discord.com/developers/applications)
- Create a publicly accessible endpoint secured with TLS that the bot will listen for connections on.

TLS is required as this bot implementation uses the interaction endpoint, which discord requires to be secured with TLS. The recommended way to do this is via a reverse proxy like nginx. If you are using Tailscale (which you likely are), you can expose this bot via a [funnel](https://tailscale.com/kb/1223/funnel). The bot supports the `tsnet` library which allows for having this application connect directly as a machine to your tailnet without needing to run some sidecar container. Configuring tsnet is covered in [the tsnet doc](/docs/tsnet.md).

## Creating the bot

To create the bot, you must first create a bot on the [discord developer portal](https://discord.com/developers/applications). You can do this by following the following steps.

1. Go to the [discord developer portal](https://discord.com/developers/applications)
2. Click "New Application"
3. Give your application a name and click "Create"
4. Note the Application ID and Public Key, you will need this later
5. Under `Interactions Endpoint URL`, enter the URL of your bot. Remember, this must be secured with TLS.
6. Click "Bot" in the left hand menu
7. Click the "Reset Token" button and note the token, you will need this later
8. Unselect "Public Bot" and click "Save Changes"
9. Under "Bot Permission", give it the `Send Messages` permission.
10. Go to "OAuth2" in the left hand menu, then to "OAuth2 URL Generator"
11. Under "Scopes", select `bot`
12. Under "Bot Permissions", select `Send Messages`
13. Copy the URL and visit it in your browser. Select the server you want to add the bot to and click "Authorize"

This should get your bot installed on your server.

## Configuration

The bot is configured via a YAML file. The following is an example configuration file.

```yaml
server:
  address: :8032 # The address to run your server on
api:
  address: "http://localhost:7032" # the address of your Texit API
  apiKey: changeme # the API key for your Texit API
discord:
  authorized:
    - "155875705417236480" # the user and role IDs of the people who can use the bot
  guildIds:
    - "948052547795574794" # the guild IDs of the servers the bot can be used on. This is used to update commands at launch. If this isn't set, the bot will update the global scope of commands.
  applicationId: "123123123123123123" # the application ID of your bot
  publicKey: "sdkjgheoirugheirnfjenfjebiyfgewpdfm[oemfpiuebrigfu]" # the public key of your bot
  token: "skdjfhkjsdfjhsgfjhgsdjhfgsjkdfgkjshgdfjhgfjheferf" # the token of your bot
```

Fill out the file with data you collected in the previous steps. If you want to ensure that you scope your bot to only be used by you in your guild, you'll need to get your userId and guildId. You can do this by enabling developer mode in discord and right clicking on your user or guild and clicking "Copy ID".

## Installation

The bot is a binary that can be grabbed from the releases, but the recommended way is to run it via docker. You can find an example `docker-compose.yaml` and config file in the [examples directory](/examples//discord/docker/). This guide will only cover docker installation.

The following is an example `docker-compose.yaml` file that will run the bot.

```yaml
services:
  discord_bot:
    image: ghcr.io/awlsring/texit-discord:latest
    ports:
      - 8032:8032
    volumes:
      # This is the path to your config file
      - ./config.yaml:/etc/texit-discord/config.yaml
      # If you are using TS net, you'll want to mount the data directory to persist your state data.
      - ./data:/var/lib/texit-discord
```

You only need to specify the config file, which needs to map to whereever you saved your config.

## Running the bot

You can run the bot via the docker-compose file with the following command.

```bash
$ docker-compose up
```

This will stand up the bot. Logs should be flooding your terminal.

Once the bot is up, go to you discord server and type `/` in a channel. You should see the bot's commands pop up. If you don't, you may need to wait a few minutes for discord to update the commands.

Before you run a command, ensure you have set the interaction endpoint in the discord developer portal and have done whatever reverse proxy / TLS setup you need to do to make the bot accessible.

Once you have done that, you can sanity check that it stood up by running the `/self-health` command.
