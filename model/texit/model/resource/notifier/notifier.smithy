$version: "2.0"

namespace awlsring.texit.api

@documentation("The name of the notifier.")
@length(min: 1, max: 36)
string NotifierName

resource Notifier {
    identifiers: {name: NotifierName}
    list: ListNotifiers
}

@documentation("The type of notifier.")
enum NotifierType {
    AWS_SNS = "aws-sns"
    MQTT = "mqtt"
    UNKNOWN = "unknown"
}

structure NotifierSummary {
    @required
    name: NotifierName

    @required
    type: NotifierType

    @documentation("The endpoint that is used to send notifications. For SNS, this is the topic arn. For MQTT, this is the broker address and topic.")
    @required
    endpoint: String
}

list NotifierSummaries {
    member: NotifierSummary
}
