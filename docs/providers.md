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
