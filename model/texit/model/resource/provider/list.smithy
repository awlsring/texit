$version: "2.0"

namespace awlsring.texit.api

@documentation("List all registered providers.")
@readonly
@http(method: "GET", uri: "/provider", code: 200)
operation ListProviders {
    input := {}

    output := {
        @required
        summaries: ProviderSummaries
    }
}
