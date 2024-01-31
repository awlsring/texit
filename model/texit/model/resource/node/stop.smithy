$version: "2.0"

namespace awlsring.texit.api

@documentation("Stops the target node.")
@http(method: "POST", uri: "/node/{identifier}/stop", code: 200)
operation StopNode {
    input := {
        @httpLabel
        @required
        identifier: NodeIdentifier
    }

    output := {
        @required
        success: Boolean
    }

    errors: [
        ResourceNotFoundError
    ]
}
