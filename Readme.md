# felixwie services

This mono repository contains the code for all my microservices.
Most services are written in go, some are server side rendered React projects build with Remix.

## services

Read more about each service in the readme file in its directory.

### api-gateway

Web entry point for all backend-only services.

### auth

The auth microservice handles authentication for all gRPC services with Azure Active Directory

### meal

tbd

## contributing

The monorepo management relies on common makefile targets and a yeoman generator. For more information read the [contributing guidelines](contributing.md)
