$version: "2.0"

namespace awlsring.texit.api

@documentation("Provision a node on the specified provider in a given location on the specified tailnet.")
@http(method: "POST", uri: "/node", code: 200)
operation ProvisionNode {
    input := {
        @required
        provider: ProviderName

        @required
        location: ProviderLocation

        @required
        tailnet: TailnetName

        ephemeral: Boolean

        size: NodeSize
    }

    output := {
        @required
        execution: ExecutionIdentifier
    }

    errors: [
        ResourceNotFoundError
        InvalidInputError
    ]
}
