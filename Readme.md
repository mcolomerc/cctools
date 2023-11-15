
# Apache Kafka Migration Tools

![GitHub release](https://img.shields.io/github/v/release/mcolomerc/cctools)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mcolomerc/cctools)
![GitHub](https://img.shields.io/github/license/mcolomerc/cctools)

<img src="./docs/src/.vuepress/public/logo.png" width="100"> 

[cctools](https://mcolomerc.github.io/cctools/) is command Line tool for helping on Kafka migrations to Confluent Cloud or Confluent Platform.

This CLI uses Kafka client and REST APIs to extract and export all the resources from the Source cluster in order to replicate them on the target cluster. It was tested with Confluent Platform and Confluent Cloud clusters.

It allows to export resources into different formats, that could be used as input for different tools like Confluent Cloud, Terraform, Confluent For Kubernetes or any other tool.

It provides a `copy` command to replicate resources configuration between clusters.

## Installation

Go to [Releases](https://github.com/mcolomerc/cctools/releases) and Download your OS distribution.

## Usage

Commands:

- `copy` Copy resources between clusters
  
- `export` Export resources from a cluster

- `import` Import resources to a cluster

Go to [Docs](https://mcolomerc.github.io/cctools/) for more information.

---

# Sources

## Execute

Export everything to all the formats available:

* `go run main.go  export  --config config_cloud.yml`

* Export everything to JSON format:

`go run main.go  export --output json --config config_cloud.yml`

* Export everything to JSON & YAML format:
`go run main.go  export --output json,yaml --config config_cloud.yml`   (Not supported)

* Export Topics:

`go run main.go  export topics --output json --config config_cloud.yml`

* Export ACLs:

`go run main.go  export acls --output json --config config_cloud.yml`

* Export Schema Registry:

`go run main.go export schemas --output yaml --config config_cloud.yml`

## DEBUG

Enable Debug mode:`LOG=DEBUG` for extra logging.

## Releaser

<https://goreleaser.com/install/>

```brew install goreleaser```

```goreleaser build --snapshot --rm-dist```

### Binary

MacOS:

```./dist/cctools_darwin_amd64_v1/cctools export --config config.yml --output json```

### Docs

[VuePress](https://v2.vuepress.vuejs.org/) site: `docs`