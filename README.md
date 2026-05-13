# Temporal Cloud Resource Provider

The Temporal Cloud Resource Provider lets you manage [Temporal Cloud](https://temporal.io/cloud) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @nellisauction/pulumi-temporalcloud
```

or `yarn`:

```bash
yarn add @nellisauction/pulumi-temporalcloud
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_temporalcloud
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/nellisauction/pulumi-temporalcloud/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.TemporalCloud
```

## Configuration

The following configuration points are available for the `temporalcloud` provider:

- `temporalcloud:apiKey` (environment: `TEMPORAL_CLOUD_API_KEY`) - a Temporal Cloud API key
- `temporalcloud:endpoint` (environment: `TEMPORAL_CLOUD_ENDPOINT`) - the Temporal Cloud API endpoint (default: `saas-api.tmprl.cloud:443`)
- `temporalcloud:allowInsecure` (environment: `TEMPORAL_CLOUD_ALLOW_INSECURE`) - allow insecure connections (default: `false`)
- `temporalcloud:allowedAccountId` (environment: `TEMPORAL_CLOUD_ALLOWED_ACCOUNT_ID`) - restrict operations to a specific account ID

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/temporalcloud/api-docs/).
