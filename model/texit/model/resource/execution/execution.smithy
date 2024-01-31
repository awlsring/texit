$version: "2.0"

namespace awlsring.texit.api

@documentation("A node's identifier.")
@length(min: 40, max: 60)
string ExecutionIdentifier

resource Execution {
    identifiers: {identifier: ExecutionIdentifier}
    read: GetExecution
}

@documentation("The status of an execution.")
enum ExecutionStatus {
    PENDING = "pending"
    RUNNING = "running"
    COMPLETED = "completed"
    FAILED = "failed"
    UNKNOWN = "unknown"
}

@documentation("The name of a workflow.")
enum WorkflowName {
    PROVISION_NODE = "provision-node"
    DEPROVISION_NODE = "deprovision-node"
    UNKNOWN = "unknown"
}

structure ExecutionSummary {
    @required
    status: ExecutionStatus

    @required
    workflow: WorkflowName

    @required
    startedAt: Timestamp

    endedAt: Timestamp

    result: String
}
