# Core Deployment

This folder has the following files:

- `ops`: A shell script to run docker images that initialize a `blockstack-core` node's `~/.blockstack-server` directory with the `fast_sync` directive
- `blockstack-server.ini`: A config file for the docker images in `ops`
- `docker-compose.yaml`: Runs 4 `blockstack-core` nodes, each on a separate host port.

This allows for running tightly packed `blockstack-core` nodes on multi-core servers.

Test core node connectivity with the following curl call:

```
curl -L -XPOST -H "Content-Type: application/xml" {addresss}:{port}/RPC2 -d '<?xml version="1.0"?><methodCall><methodName>getinfo</methodName><params></params></methodCall>'
```
