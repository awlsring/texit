$version: "2.0"

namespace awlsring.texit.api

@documentation("Get the summary of a provider.")
@readonly
@http(method: "GET", uri: "/tailnet/{name}", code: 200)
operation DescribeTailnet {
    input := {
        @httpLabel
        @required
        name: TailnetName
    }

    output := {
        @required
        summary: TailnetSummary
    }

    errors: [
        InvalidInputError
        ResourceNotFoundError
    ]
}
