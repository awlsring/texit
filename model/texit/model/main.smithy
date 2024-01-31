$version: "2.0"

namespace awlsring.texit.api

use aws.protocols#restJson1

@restJson1
@title("Texit")
@httpApiKeyAuth(name: "X-Api-Key", in: "header")
@paginated(inputToken: "nextToken", outputToken: "nextToken", pageSize: "pageSize")
service Texit {
    version: "2024-01-31"
    operations: [Health]
    resources: [Provider, Tailnet, Execution, Node]
}
