# Notifier

A "notifier" is some system that will be called on the completion of a texit workflow to notify of the result. Texit supports having numerous notifiers and if multiple are configured all will be called on the completion of a workflow. Unlike providers and tailnets, notifiers are not required to run Texit.

Notifiers are configured in the `notifiers` block of your texit configuration file. The following sections will detail different types of notifiers and properties that must be set to utilize them.

## MQTT

The MQTT notifier allows for sending a message to a MQTT broker on the completion of a texit workflow. Its type is `mqtt`.

You must have a pre-existing MQTT broker to use this notifier. Texit expects to communicate with a broker using MQTT 3.1.

### Configuration

To include a MQTT notifier for your texit, you must include the following extra fieilds in your `notifiers` configuration block config file.

- **Type**: Indicate the type of notifier. This will be `mqtt` for this notifier.
- **Topic**: The topic to publish the message to. This is required.
- **Broker**: The address of the MQTT broker to connect to. This will look something like `tcp://10.0.1.2:1883` This is required.
- **Username**: The username to authenticate with the MQTT broker. This is optional.
- **Password**: The password to authenticate with the MQTT broker. This is optional.

#### Example

```yaml
notifiers:
  - type: mqtt
    broker: "tcp://10.0.1.2:1883"
    topic: "texit-workflow"
```

## AWS-SNS

The AWS-SNS notifier allows for sending a message to a SNS topic on the completion of a texit workflow. Its type is `sns`.

### Configuration

To include a SNS notifier for your texit, you must include the following extra fieilds in your `notifiers` configuration block config file.

- **Type**: Indicate the type of notifier. This will be `sns` for this notifier.
- **Topic**: The topic arn of the topic to publish to. This is required. This can also be set with an env variable of `SNS_NOTIFIER_ARN`.
- **Region**: The region of the SNS topic to publish to. This is required.
- **AccessKey**: This is an access key for the IAM user texit authenticates as when calling SNS. This can also be set with an env variable of `SNS_AWS_ACCESS_KEY_ID`. This is required.
- **SecretKey**: This is the secrey key for the IAM user texit authenticates as when calling SNS. This can also be set with an env variable of `SNS_AWS_SECRET_ACCESS_KEY`. This is required.

#### Example

```yaml
notifiers:
  - type: sns
    accessKey: "XXXXXXXXXXXXXX"
    secretKey: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    region: "us-west-2"
    topic: "arn:aws:sns:us-west-2:0101010101011:MyTexitTopic"
```
