$version: "2.0"

namespace awlsring.texit.api

@documentation("Lists all known nodes.")
@readonly
@http(method: "GET", uri: "/node", code: 200)
operation ListNodes {
    input := {}

    output := {
        @required
        summaries: NodeSummaries
    }
}
