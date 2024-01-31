$version: "2.0"

namespace awlsring.texit.api

@documentation("Get the summary of an node.")
@readonly
@http(method: "GET", uri: "/node/{identifier}", code: 200)
operation DescribeNode {
    input := {
        @httpLabel
        @required
        identifier: NodeIdentifier
    }

    output := {
        @required
        summary: NodeSummary
    }

    errors: [
        ResourceNotFoundError
    ]
}
