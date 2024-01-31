$version: "2.0"

namespace awlsring.texit.api

@documentation("Deprovision the target node.")
@idempotent
@http(method: "DELETE", uri: "/node/{identifier}", code: 200)
operation DeprovisionNode {
    input := {
        @httpLabel
        @required
        identifier: NodeIdentifier
    }

    output := {
        @required
        execution: ExecutionIdentifier
    }

    errors: [
        ResourceNotFoundError
    ]
}
