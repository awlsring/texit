$version: "2.0"

namespace awlsring.texit.api

@documentation("Get the summary of a provider.")
@readonly
@http(method: "GET", uri: "/provider/{name}", code: 200)
operation DescribeProvider {
    input := {
        @httpLabel
        @required
        name: ProviderName
    }

    output := {
        @required
        summary: ProviderSummary
    }

    errors: [
        ResourceNotFoundError
    ]
}
