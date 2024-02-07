$version: "2.0"

namespace awlsring.texit.api

@documentation("Starts the target node.")
@http(method: "POST", uri: "/node/{identifier}/start", code: 200)
operation StartNode {
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
        InvalidInputError
        ResourceNotFoundError
    ]
}
