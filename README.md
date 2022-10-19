# Songbird & Coston 1

Node implementation for the [Flare](https://flare.network) network.

## Running containers

*The fast and simple way of using go-songbird*

Public container images are hosted on Docker HUB and Github Packages;
```
docker.io/flarenetwork/go-songbird
hgcr.io/flare-foundation/go-songbird
```

### Container builds in CI

CI builds on each:
- push on `main` branch, pushes image tagged as "latest"
- creation of a tag, pushes images tagged as the tag itself

Builds: \
two images, `go-songbird:<TAG>` one with `leveldb` and `go-songbird:<TAG>-rocksdb` with RocksDB builtin

### Build arguments

| Argument name | Default value | description |
|---|---|---|
| `DB_TYPE` | `leveldb` | if `rocksdb` the image will be built with rocksdb support; ref [docs.avax.network](https://docs.avax.network/nodes/maintain/avalanchego-config-flags#database) |


### Runtime environment variables

| Varible name | Default value | description |
|---|---|---|
| `HTTP_HOST` | `0.0.0.0` | Should always be `0.0.0.0` as it's a container |
| `HTTP_PORT` | `9650` | |
| `STAKING_PORT` | `9651` | |
| `PUBLIC_IP` | ` ` | can be autoconfigured by having `AUTOCONFIGURE_PUBLIC_IP` enabled |
| `DB_TYPE` | `leveldb` | One of `leveldb \| rocksdb \| memdb \| memdb`. Rocksdb can only be used with images whose tags end with `-rocksdb`. |
| `DB_DIR` | `/app/db` | |
| `BOOTSTRAP_IPS` | ` ` | [--bootstrap-ids-string](https://docs.avax.network/nodes/maintain/avalanchego-config-flags#--bootstrap-ids-string), can be autoconfigured by enabling `AUTOCONFIGURE_BOOTSTRAP` |
| `BOOTSTRAP_IDS` | ` ` | [--bootstrap-ips-string](https://docs.avax.network/nodes/maintain/avalanchego-config-flags#--bootstrap-ips-string), can be autoconfigured by enabling `AUTOCONFIGURE_BOOTSTRAP` |
| `CHAIN_CONFIG_DIR` | `/app/conf` | |
| `LOG_DIR` | `/app/logs` | |
| `LOG_LEVEL` | `info` | |
| `NETWORK_ID` | `coston` | One of `songbird \| coston` Define the [target network](https://docs.flare.network/dev/reference/network-configs/) to work with |
| `AUTOCONFIGURE_PUBLIC_IP` | `1` | Autoconfigure PUBLIC_IP, skipped if PUBLIC_IP is set |
| `AUTOCONFIGURE_BOOTSTRAP` | `1` | Enables auto-fetch of [--bootstrap-ids-string](https://docs.avax.network/nodes/maintain/avalanchego-config-flags#--bootstrap-ids-string) and [--bootstrap-ips-string](https://docs.avax.network/nodes/maintain/avalanchego-config-flags#--bootstrap-ips-string) values from `AUTOCONFIGURE_BOOTSTRAP_ENDPOINT` |
| `AUTOCONFIGURE_BOOTSTRAP_ENDPOINT` | `https://coston.flare.network/ext/info` | Endpoint used for [bootstrapping](https://docs.avax.network/nodes/maintain/avalanchego-config-flags#bootstrapping) info fetch |
| `EXTRA_ARGUMENTS` | ` ` | Extra arguments passed to flare binary when running it from `entrypoint.sh` |

---


## Installation

Flare uses a relatively lightweight consensus protocol, so the minimum computer requirements are modest.
Note that as network usage increases, hardware requirements may change.

The minimum recommended hardware specification for nodes connected to Mainnet is:

- CPU: Equivalent of 8 vCPU
- RAM: 16 GiB
- Storage: 2.5TB for a full archive node 1TB for pruning
- OS: Ubuntu 18.04/20.04 or macOS >= 10.15 (Catalina)
- Network: Reliable IPv4 or IPv6 network connection, with an open public port.

If you plan to build Flare from source, you will also need the following software:

- [Go](https://golang.org/doc/install) version >= 1.16.8
- [gcc](https://gcc.gnu.org/)
- g++

### Native Install

Clone the Flare repository:

```sh
git clone https://github.com/flare-foundation/go-songbird.git
cd go-songbird/avalanchego
```

This will clone and checkout to `master` branch. 

Please build and use the latest tag `0.6.4`

### Building the Flare Executable

Build Flare using the build script:

```sh
./scripts/build.sh
```

The service binary is named `flare` and is in the `build` directory.


## Running the Flare binary

### Connecting to Coston

To connect to the Coston test network, run:

```sh
./build/flare --network-id=coston \
  --bootstrap-ips="$(curl -m 10 -sX POST --data '{ "jsonrpc":"2.0", "id":1, "method":"info.getNodeIP" }' -H 'content-type:application/json;' https://coston.flare.network/ext/info | jq -r ".result.ip")" \
  --bootstrap-ids="$(curl -m 10 -sX POST --data '{ "jsonrpc":"2.0", "id":1, "method":"info.getNodeID" }' -H 'content-type:application/json;' https://coston.flare.network/ext/info | jq -r ".result.nodeID")"
```

You should see some _fire_ ASCII art and log messages.

You can use `Ctrl+C` to kill the node.

If you want your node's API to be reachable, you have to add the `--http-host=<ip_address>` flag to the command line.

### Connecting to Songbird

To connect to the Songbird canary network, run:

```sh
./build/flare --network-id=songbird \
  --bootstrap-ips="$(curl -m 10 -sX POST --data '{ "jsonrpc":"2.0", "id":1, "method":"info.getNodeIP" }' -H 'content-type:application/json;' https://songbird.flare.network/ext/info | jq -r ".result.ip")" \
  --bootstrap-ids="$(curl -m 10 -sX POST --data '{ "jsonrpc":"2.0", "id":1, "method":"info.getNodeID" }' -H 'content-type:application/json;' https://songbird.flare.network/ext/info | jq -r ".result.nodeID")"
```

You should see some _fire_ ASCII art and log messages.

You can use `Ctrl+C` to kill the node.

If you want your node's API to be reachable, you have to add the `--http-host=<ip_address>` flag to the command line.

Please note that you currently need to be whitelisted in order to connect to the Songbird network.

### Pruning & APIs

The configuration for the chain is loaded from a configuration file, located at `{chain-config-dir}/C/config.json`.

Here are the most relevant default settings:

```json
{
  "snowman-api-enabled": false,
  "coreth-admin-api-enabled": false,
  "eth-apis": [
    "public-eth",
    "public-eth-filter",
    "net",
    "web3",
    "internal-public-eth",
    "internal-public-blockchain",
    "internal-public-transaction-pool"
  ],
  "rpc-gas-cap": 50000000,
  "rpc-tx-fee-cap": 100,
  "pruning-enabled": true,
  "local-txs-enabled": false,
  "api-max-duration": 0,
  "api-max-blocks-per-request": 0,
  "allow-unfinalized-queries": false,
  "allow-unprotected-txs": false,
  "remote-tx-gossip-only-enabled": false,
  "log-level": "info"
}
```

You can refer to the original [Avalanche documentation](https://docs.avax.network/build/references/avalanchego-config-flags/#c-chain-configs) for a full list of all settings and a detailed description.

The directory for configuration files defaults to `$HOME/.flare/configs/chains` and can be changed using the `--chain-config-dir` flag.

In order to disable pruning and run a full archival node, `pruning-enabled` should be set to `false`.

The various node APIs can also be enabled and disabled by setting the respective parameters.

### Launching Flare locally

In order to run a local network, the validator set needs to be defined locally.
This can be done by setting the validator set in a environment variable.

You can use `./scripts/launch_localnet.sh` as an easy way to spin up a 5-node local network.
All funds are controlled by the private key under `/.scripts/keys/6b0dd034a2fd67b932f10e3dba1d2bbd39348695.json`.

## Generating Code

Flare uses multiple tools to generate boilerplate code.

### Running protobuf codegen

To regenerate the protobuf go code, run `scripts/protobuf_codegen.sh` from the root of the repo.

This should only be necessary when upgrading protobuf versions or modifying .proto definition files.

To use this script, you must have [buf](https://docs.buf.build/installation) (v1.0.0-rc12), protoc-gen-go (v1.27.1) and protoc-gen-go-grpc (v1.2.0) installed.

To install the buf dependencies:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
```

If you have not already, you may need to add `$GOPATH/bin` to your `$PATH`:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

If you extract buf to ~/software/buf/bin, the following should work:

```sh
export PATH=$PATH:~/software/buf/bin/:~/go/bin
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/protobuf/cmd/protoc-gen-go-grpc
scripts/protobuf_codegen.sh
```

For more information, refer to the [GRPC Golang Quick Start Guide](https://grpc.io/docs/languages/go/quickstart/).        |

## Security Bugs

**We and our community welcome responsible disclosures.**

If you've discovered a security vulnerability, please report it via our [contact form](https://flare.network/contact/). Valid reports will be eligible for a reward (terms and conditions apply).
