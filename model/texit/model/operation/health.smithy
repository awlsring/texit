$version: "2.0"

namespace awlsring.texit.api

@readonly
@http(method: "GET", uri: "/health", code: 200)
operation Health {
    input := {}
    output := {
        healthy: Boolean
    }
}
