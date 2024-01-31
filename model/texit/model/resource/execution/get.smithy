$version: "2.0"

namespace awlsring.texit.api

@documentation("Get the summary of an execution.")
@readonly
@http(method: "GET", uri: "/execution/{identifier}", code: 200)
operation GetExecution {
    input := {
        @httpLabel
        @required
        identifier: ExecutionIdentifier
    }

    output := {
        @required
        summary: ExecutionSummary
    }

    errors: [
        ResourceNotFoundError
    ]
}
