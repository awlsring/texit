$version: "2.0"

namespace awlsring.texit.api

@documentation("The name of the provider.")
@length(min: 1, max: 36)
string ProviderName

resource Provider {
    identifiers: {name: ProviderName}
    read: DescribeProvider
    list: ListProviders
}

@documentation("The type of provider.")
enum ProviderType {
    AWS_ECS = "aws-ecs"
    UNKNOWN = "unknown"
}

@documentation("A location provided by a provider.")
string ProviderLocation

@documentation("The identifier of the node resource in the provider.")
string ProviderNodeIdentifier

structure ProviderSummary {
    @required
    name: ProviderName

    @required
    type: ProviderType
}

list ProviderSummaries {
    member: ProviderSummary
}
