$version: "2.0"

namespace awlsring.texit.api

@documentation("List all registered tailnets.")
@readonly
@http(method: "GET", uri: "/tailnet", code: 200)
operation ListTailnets {
    input := {}

    output := {
        @required
        summaries: TailnetSummaries
    }
}
