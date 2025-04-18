# v1.3.1 (2023-08-01)

* No change notes available for this release.

# v1.3.0 (2023-07-31)

* **Feature**: Adds support for smithy-modeled endpoint resolution. A new rules-based endpoint resolution will be added to the SDK which will supercede and deprecate existing endpoint resolution. Specifically, EndpointResolver will be deprecated while BaseEndpoint and EndpointResolverV2 will take its place. For more information, please see the Endpoints section in our Developer Guide.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.14 (2023-07-28)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.13 (2023-07-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.12 (2023-06-15)

* No change notes available for this release.

# v1.2.11 (2023-06-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.10 (2023-05-04)

* No change notes available for this release.

# v1.2.9 (2023-04-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.8 (2023-04-10)

* No change notes available for this release.

# v1.2.7 (2023-04-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.6 (2023-03-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.5 (2023-03-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.4 (2023-02-22)

* **Bug Fix**: Prevent nil pointer dereference when retrieving error codes.

# v1.2.3 (2023-02-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.2 (2023-02-15)

* **Announcement**: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR #2012 tracked in issue #1910.
* **Bug Fix**: Correct error type parsing for restJson services.

# v1.2.1 (2023-02-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.0 (2023-01-05)

* **Feature**: Add `ErrorCodeOverride` field to all error structs (aws/smithy-go#401).

# v1.1.4 (2022-12-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.3 (2022-12-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.2 (2022-10-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.1 (2022-10-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.0 (2022-10-19)

* **Feature**: This release adds the RegisterSlackWorkspaceForOrganization API. You can use the API to register a Slack workspace for an AWS account that is part of an organization.

# v1.0.5 (2022-09-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.4 (2022-09-14)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.3 (2022-09-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.2 (2022-08-31)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.1 (2022-08-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.0 (2022-08-22)

* **Release**: New AWS service client module
* **Feature**: This is the initial SDK release for the AWS Support App in Slack.

