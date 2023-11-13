--- 
title: Guide
description: Guide page 
tags:
  - guide
  - installation
prev: ../config/README.md
next: ../commands/README.md
---
 
# Command Line ![GitHub release](https://img.shields.io/github/v/release/mcolomerc/cctools)

 ```cctools``` is command Line tool for helping on Kafka operations or migrations between different Kafka platforms. Cluster migrations are not a trivial task, and this tool aims to help on this process.

This CLI uses Kafka client and REST APIs to extract and export all the resources from the Source cluster in order to replicate them on the target cluster. It was tested with Confluent Platform and Confluent Cloud clusters.

It allows to `export` resources into different formats, that could be used as input for different tools like [Confluent Cloud](https://www.confluent.io/lp/confluent-cloud), [Terraform](https://registry.terraform.io/providers/confluentinc/confluent/latest), [Confluent For Kubernetes](https://docs.confluent.io/operator/current/overview.html) or any other tool.

It provides a `copy` command to replicate cluster metadata, like topics or consumer groups, between clusters.

It is possible to `import` resources to a cluster.

## Installation

Go to [Releases](https://github.com/mcolomerc/cctools/releases) and Download your OS distribution.

Last Release: ![GitHub release](https://img.shields.io/github/v/release/mcolomerc/cctools)

## Example Usage

Export to JSON files all the Topics metadata from a Source Cluster. See [Topic Exporters](../commands/export/topics.md) documentation for using other exporters.

Create a `config.yaml` file:

```sh:no-line-numbers
touch config.yaml
```

- Add the Source Cluster connection configuration, the `source` tag requires:
  
- Cluster: `kafka` section contains the source Kafka cluster connection configuration.
  - `bootstrapServer`: Source Cluster bootstrap server.
  - `clientProps`: Kafka client properties map.

Example for SASL_SSL (*config.yaml*):

```yaml
source: 
  kafka:
    bootstrapServer: <bootstrap_server>
    clientProps:  
      - ssl.ca.location: "<path>/cacerts.pem" 
      - sasl.mechanisms: PLAIN
      - security.protocol: SASL_SSL
      - sasl.username: <admin_user>
      - sasl.password: <admin_password>
```

- Add the output path for the JSON files, all the Topics metadata will be exported to `<export.output.path>/topics/json`, one file per Topic.

```yaml
export:
  output: <path>
```

- Filtering Topics. Exclude.

```yaml
export:
  output: <path>
  topics:
    exclude: <topic_name_substring>  #_confluent
```

- Run the command:

```sh:no-line-numbers
cctools export topics --config config.yaml --output json
```

- Output folder. One JSON file per Topic exported.

JSON Topic example:

```json
{
 "name": "demo.topic",
 "partitions": 4,
 "replicationFactor": 3,
 "minIsr": "2",
 "retentionTime": "604800000",
 "configs": [
  {
   "name": "confluent.stray.log.max.deletions.per.run",
   "value": "72"
  },
  {
   "name": "confluent.stray.log.delete.delay.ms",
   "value": "604800000"
  },
  ...
 ],
  "acls": [
  {
   "principal": "User:test",
   "host": "*",
   "operation": "ALL",
   "permission": "ALLOW",
   "resourceType": "TOPIC",
   "resourceName": "demo.topic",
   "patternType": "LITERAL"
  }
 ]
```

## Configuration

See documentation for all the [Configuration](../config/README.md)
 
## Commands

### Copy

 ```sh:no-line-numbers
 cctools copy
 ```
 
 A command to replicate cluster metadata, like topics or consumer groups, between clusters.

Usage: [Copy](../commands/copy.md)

See documentation for all the [Commands](../commands/README.md)

### Export

 ```sh:no-line-numbers
 cctools export
 ```

 A command to export cluster metadata to different formats with a single command.

Usage: [Export](../commands/export.md)

See documentation for all the [Commands](../commands/README.md)

### Import

 ```sh:no-line-numbers
 cctools import
 ```

 A command to import metadata.

Usage: [Import](../commands/import.md)

See documentation for all the [Commands](../commands/README.md)