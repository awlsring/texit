$version: "2.0"

namespace awlsring.texit.api

@documentation("Get the status of an node.")
@readonly
@http(method: "GET", uri: "/node/{identifier}/status", code: 200)
operation GetNodeStatus {
    input := {
        @httpLabel
        @required
        identifier: NodeIdentifier
    }

    output := {
        @required
        status: NodeStatus
    }

    errors: [
        ResourceNotFoundError
    ]
}
