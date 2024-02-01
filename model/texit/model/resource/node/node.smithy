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
    operations: [StartNode, StopNode, GetNodeStatus]
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

    @required
    location: ProviderLocation

    @required
    providerNodeIdentifier: ProviderNodeIdentifier

    @required
    tailnet: TailnetName

    @required
    tailnetDeviceName: TailnetDeviceName

    @required
    TailnetDeviceIdentifier: TailnetDeviceIdentifier

    @required
    @documentation("If a node is ephemeral.")
    ephemeral: Boolean

    @required
    @documentation("When a node was created.")
    created: Timestamp

    @required
    @documentation("When a node was last updated.")
    updated: Timestamp
}

list NodeSummaries {
    member: NodeSummary
}
