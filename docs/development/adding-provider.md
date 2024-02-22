# Adding a New Provider

This doc is to detail the process of adding a new provider. Its outlines what additions to which parts of the codebase are necessary to update.

Highlevel overview of the process:

- Pick a provider type name. This should be short and unique amoungst providers.
- Add this provider name to the SmithyModel enum `ProviderType` in `model/texit/model/resource/provider/provider.smithy`.
- Add this provider name to the Go enum `provider.Type` in `internal/pkg/domain/provider/type.go`. Ensure it has entries in the `String` and `TypeFromString` methods.
- Define configuration that must be passed by a user in the config file. Add these fields to the provider config struct in `internal/app/api/config/provider.go`. Create a new function to validate this provider to ensure the required fields are present.
- Add a section detailing this provider to the `providers.md` doc in `docs/providers.md`. This should include the required fields and a brief description of the provider.
- Write a new adapter in `internal/app/api/adapters/secondary/gateway/platform/<mod for new provider>`. This should contain a struct that implements the `gateway.Platform` interface.
- Add the location options to the Discord autocomplete in `internal/app/ui/adapters/primary/discord/handler/location_auto_complete.go`.
- Add the provider to the `translateProviderType` method in `internal/app/ui/adapters/secondary/gateway/api/conversion.go`.
- Add the provider to the `TranslateProviderType` method in `internal/app/api/adapters/primary/ogen/conversion/provider.go`.
- Add this provider to the `LoadProviderGateways` method in `internal/app/api/setup/gateways.go`.

Example commit adding the `linode` provider: https://github.com/awlsring/texit/commit/3da91eac9d152937b09e0ed899bd69d03bb8b585#diff-c36280bd75640ab0b0ea6b095c47b0b74dd2493beee15846d6cf5ab72f4104d1
