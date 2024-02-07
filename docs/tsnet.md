# TSNet Support

Both the Texit API and the Texit Discord bot support the `tsnet` library. This library allows for having this application connect directly as a machine to your tailnet without needing to run some additional process to connect to the tailnet.

## Configuration

Both application follows a standard configuration under their `server` configuration block. The following is an example configuration:

```yaml
server:
  address: :443
  tailnet:
    authkey: tskey-auth-xxxxxxx-xxxxxxxxxxxxxxxx
    mode: funnel
    hostname: host
    state: /var/lib/texit
```

This configuration allows for specifying the following fields:

- **Authkey**: This is the authkey that allows the application to connect to the tailnet without approval. This is required.
- **Mode**: This is the mode that the application will use to connect to the tailnet. This can be `funnel` or `tls`. `standard`. _Only Tailscale supports `funnel` and `tls` mode_.
- **Hostname**: This is the hostname that the application will use to connect to the tailnet. This is optional and by default is set by the application.
- **State**: This is the state directory that the application will use to store its state. This is optional and by default is set by the application.
- **ControlUrl**: This is the control url that the application will use to connect to the tailnet. This only needs to be set if you are using Headscale.
