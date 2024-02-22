# Providers

A "provider" is some platform that texit supports that can support running exit nodes. Texit supports running numerous providers. All providers must be specified in your texit configuration file under the `providers` block.

Some providers will have unique fields that must be set. All providers will need to specify at least the following...

- **Name**: This is the name you want to call your provider and how you will target it when calling texit apis. This must be unique amoungst all your providers
- **Type**: This is the type of provider you are configuring. This must match the type of provider you are specifying.

Further sections will detail all supported providers and unique fields that must be set when configuring them.

## AWS-ECS

AWS-ECS provider allows for launching exit nodes as Fargate containers on ECS. Its type is `aws-ecs`.

You must create an IAM user to use this provider. You can get instructions for how to set up this provider with texit cli running the following command.

```
$ texit provider init -t aws-ecs
```

### Configuration

To include a AWS-ECS provider for your texit, you must include the following extra fieilds in your `provider` configuration block config file.

- **AccessKey**: This is an access key for the IAM user texit authenticates as when calling AWS apis. This can also be set with an env variable of `<provider-name>_AWS_ACCESS_KEY_ID`.
- **SecretKey**: This is the secrey key for the IAM user texit authenticates as when calling AWS apis. This can also be set with an env variable of `<provider-name>_AWS_SECRET_ACCESS_KEY`.

#### Example

```yaml
providers:
  - type: aws-ecs
    accessKey: "XXXXXXXXXXXX"
    secretKey: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    name: "my-ecs-provider"
```

## AWS-EC2

AWS-EC2 provider allows for launching exit nodes as ec2 instances. Its type is `aws-ec2`.

This provider launches exit nodes with the instance type `t4g.nano`.

You must create an IAM user to use this provider. You can get instructions for how to set up this provider with texit cli running the following command.

```
$ texit provider init -t aws-ec2
```

### Configuration

To include a AWS-EC2 provider for your texit, you must include the following extra fieilds in your `provider` configuration block config file.

- **AccessKey**: This is an access key for the IAM user texit authenticates as when calling AWS apis. This can also be set with an env variable of `<provider-name>_AWS_ACCESS_KEY_ID`.
- **SecretKey**: This is the secrey key for the IAM user texit authenticates as when calling AWS apis. This can also be set with an env variable of `<provider-name>_AWS_SECRET_ACCESS_KEY`.

#### Example

```yaml
providers:
  - type: aws-ec2
    accessKey: "XXXXXXXXXXXX"
    secretKey: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    name: "my-ec2-provider"
```

## Linode/Akamai

The Linode (Akamai) provider allows for launching exit nodes as linodes. Its type is `linode`.

This provider launches exit nodes with the linode type `g6-nanode-1` using Debian 12.

You must create a Linode API token to use this provider. This token needs to have Read/Write scopes for `Linodes` and `StackScripts`.

### Configuration

To include a Linode provider for your texit, you must include the following extra fieilds in your `provider` configuration block config file.

- **ApiKey**: This is an access token you can create on the Linode console. This can also be set with the env variable `<provider-name>_API_KEY`.

#### Example

```yaml
providers:
  - type: linode
    apiKey: "XXXXXXXXXXXX"
    name: "my-linode-provider"
```

## Hetzner

The Hetzner provider allows for launching exit nodes as linodes. Its type is `hetzner`.

This provider launches exit nodes with the shared server type `cx11` in EU locations and `cpx11` in US location. Either server will use Debian 12. This server will be allocated a public IPv4 address.

You must create a API token to use this provider. This token needs Read and Write access.

### Configuration

To include a Hetzner provider for your texit, you must include the following extra fieilds in your `provider` configuration block config file.

- **ApiKey**: This is an access token you can create on the Linode console. This can also be set with the env variable `<provider-name>_API_KEY`.

#### Example

```yaml
providers:
  - type: hetzner
    apiKey: "XXXXXXXXXXXX"
    name: "my-hetzner-provider"
```
