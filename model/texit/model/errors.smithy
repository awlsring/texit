$version: "2.0"

namespace awlsring.texit.api

@error("client")
@httpError(400)
structure InvalidInputError {
    @required
    message: String
}

@error("client")
@httpError(404)
structure ResourceNotFoundError {
    @required
    message: String
}

@error("client")
@httpError(401)
structure UnauthorizedError {
    @required
    message: String
}

@error("server")
@httpError(500)
structure InternalServerError {
    @required
    message: String
}
