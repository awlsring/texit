$version: "2.0"

namespace awlsring.texit.api

@documentation(".")
@length(min: 1, max: 100)
string TailnetName

resource Tailnet {
    identifiers: {name: TailnetName}
    read: DescribeTailnet
    list: ListTailnets
}

enum TailnetType {
    TAILSCALE = "tailscale"
    HEADSCALE = "headscale"
    UNKNOWN = "unknown"
}

@documentation("The identifier of a tailnet device.")
string TailnetDeviceIdentifier

@documentation("The name of a tailnet device.")
string TailnetDeviceName

@documentation("Summary of a tailnet.")
structure TailnetSummary {
    @required
    name: TailnetName

    @required
    type: TailnetType

    @documentation("The server address of the tailnet. This must be set for tailscale")
    address: String

    @documentation("The user Texit acts as in the tailnet.")
    user: String
}

list TailnetSummaries {
    member: TailnetSummary
}
