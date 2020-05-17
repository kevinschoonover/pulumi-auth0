# Auth0 Resource Provider

The Auth0 Resource Provider lets you manage Auth0 resources.

Modify this README to describe:

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/auth0

or `yarn`:

    $ yarn add @pulumi/auth0

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_auth0

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-auth0/sdk/go/...
    
## Configuration

The following configuration points are available:

- `auth0:domain` - (Required) Your Auth0 domain name. It can also be sourced from the `AUTH0_DOMAIN` environment variable.
- `auth0:client_id` - (Required) Your Auth0 client ID. It can also be sourced from the `AUTH0_CLIENT_ID` environment variable.
- `auth0:client_secret` - (Required) Your Auth0 client secret. It can also be sourced from the `AUTH0_CLIENT_SECRET` environment variable.
- `auth0:debug` - (Optional) Indicates whether or not to turn on debug mode.

## Reference

For detailed reference documentation, please visit [the API docs](https://pulumi.io/reference/pkg/nodejs/@pulumi/auth0/index.html).
