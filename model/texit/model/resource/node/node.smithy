$version: "2.0"

namespace awlsring.texit.api

@documentation("A node's identifier.")
@length(min: 8, max: 8)
string NodeIdentifier

resource Node {
    identifiers: {identifier: NodeIdentifier}
    read: DescribeNode
    list: ListNodes
    create: ProvisionNode
    delete: DeprovisionNode
    operations: [StartNode, StopNode]
}

@documentation("The status of a node.")
enum NodeStatus {
    STARTING = "starting"
    RUNNING = "running"
    STOPPING = "stopping"
    STOPPED = "stopped"
    UNKNOWN = "unknown"
}

structure NodeSummary {
    @required
    identifier: NodeIdentifier

    @required
    provider: ProviderName

    location: ProviderLocation

    providerNodeIdentifier: ProviderNodeIdentifier

    tailnet: TailnetName

    tailnetDeviceName: TailnetDeviceName

    TailnetDeviceIdentifier: TailnetDeviceIdentifier

    @documentation("If a node is ephemeral.")
    ephemeral: Boolean

    @documentation("When a node was created.")
    created: Timestamp

    @documentation("When a node was last updated.")
    updated: Timestamp
}

list NodeSummaries {
    member: NodeSummary
}
