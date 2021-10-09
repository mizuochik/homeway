# Homeway

A reverse proxy to home networks

## Design

### Server

#### CLI

    ./homeway-server

#### Environemnet Variables

Key | Description
--- | ---
`PROXY-PORT` | Port number listening proxy requests

#### HTTPS GET /proxy

Establish a proxy connection.

Header | Description
--- | ---
`Homeway-Forwarding-Port` | Port number forwarding data to a client
