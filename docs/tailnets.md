# Tailnets

A "tailnet" is some service that implements the Tailscale API. Texit supports creating exit nodes on multiple tailnets. All tailnets must be specified in your texit configuration file under the `tailnets` block.

All tailnets have the following configuration options.

- **Tailnet**: This is the name of your tailnet and how you will specify it when calling the texit api. Implementation is different for each tailnet type, so see the corresponding section for more details.
- **Type**: This is the type of tailnet you are configuring. This must match the type of provider you are specifying.
- **ApiKey**: This is the api key you want to authenticate as when calling the tailnet api. This can also be set with an env variable of `<tailnet-name>_TAILNET_API_KEY`. If your tailnet includes `.`, replace them with `_`in the env variable name. For example,`my.tailnet`would be`my_tailnet_TAILNET_API_KEY`.
- **ControlServer**: This is the control server that will be specified when configuring nodes. This is optional based on the provider.

Further sections will detail all supported tailnets and unique fields that must be set when configuring them.

## Tailscale

Obviously, Tailscale is supported.

To use this provider, you must be an admin on the tailnet and be able to create a new API key. You can get instructions for how to set up this provider with texit cli running the following command.

```
$ texit tailnet init -t tailscale
```

### Configuration

To include a Tailscale tailnet for your texit, you must follow the extra configurations when setting the tailnet fields in your `tailnets` block in your config file.

- **Tailnet**: This is the organization name you see on the admin panel. This is NOT your network id, which will look something like `tailssdfsdf.ts.net`.
- **ControlServer**: You dont need to set this, tailscale's is set by default

#### Example

```yaml
tailnets:
  - apiKey: "tskey-api-XXXXXX-XXXXXXXXXXX"
    tailnet: "user.provider"
    type: "tailscale"
```

## Headscale

Headscale is an open source implementation of the Tailscale API. To use this provider, you must have a Headscale server running and access to the headscale cli. You can get instructions for how to set up this provider with texit cli running the following command.

```
$ texit tailnet init -t headscale
```

### Configuration

To include a Headscale tailnet for your texit, you must follow the extra configurations when setting the tailnet fields in your `tailnets` block in your config file.

- **Tailnet**: This is just used to identify your tailnet. It needs to be unique amoungst your tailnets, but it can really be whatever you want.
- **User**: This is your headscale user. This is required for headscale.
- **ControlServer**: This is the URL for your headscale server. Something like `https://headscale.example.com`.

#### Example

```yaml
tailnets:
  - apiKey: "XXXXXX-XXXXXXXXXXX"
    tailnet: "my-tailnet"
    type: "headscale"
    user: "user"
    controlServer: "https://headscale.example.com"
```
