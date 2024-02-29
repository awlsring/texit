$version: "2.0"

namespace awlsring.texit.api

@documentation("List all registered notifiers.")
@readonly
@http(method: "GET", uri: "/notifier", code: 200)
operation ListNotifiers {
    input := {}

    output := {
        @required
        summaries: NotifierSummaries
    }
}
