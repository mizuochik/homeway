# Homeway

A reverse proxy to home networks

## Design

### Server

#### CLI

    $ homeway-server

#### Environemnet Variables

Key | Description
--- | ---
`PROXY_PORT` | Port number listening proxy requests

#### HTTPS GET /proxy

Establish a proxy connection.

Header | Description
--- | ---
`Homeway-Public-Port` | Public port number of an application

### Client

#### CLI

    $ homeway-client

#### Environment Variables

Key | Description
--- | ---
`SERVER_URL` | Homeway server URL
`LOCAL_PORT` | Local port number of an application
`PUBLIC_PORT` | Public port number of an application
