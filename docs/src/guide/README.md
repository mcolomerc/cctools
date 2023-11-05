--- 
title: Guide
description: Guide page
---

# Command Line

 <Badge type="tip" text="v1.0.31" vertical="midle" />

 ```cctools``` is command Line tool for helping on Kafka operations or migrations between different Kafka platforms. Cluster migrations are not a trivial task, and this tool aims to help on this process.

This CLI uses Kafka client and REST APIs to extract and export all the resources from the Source cluster in order to replicate them on the target cluster. It was tested with Confluent Platform and Confluent Cloud clusters.

It allows to `export` resources into different formats, that could be used as input for different tools like Confluent Cloud, Terraform, Confluent For Kubernetes or any other tool.

It provides a `copy` command to replicate cluster metadata, like topics or consumer groups, between clusters. 

## Installation

Go to [Releases](https://github.com/mcolomerc/cctools/releases) and Download your OS distribution.

[Last Release](https://github.com/mcolomerc/cctools/releases/tag/v1.0.32)

## Configuration

See documentation for all the [Configuration](../config/README.md)
 
## Commands

### Copy

 ```sh
 cctools copy
 ```
 
 A command to replicate cluster metadata, like topics or consumer groups, between clusters.

Usage: [Copy](../commands/copy.md)

### Export

 ```sh
 cctools export
 ``` 
 
 A command to export cluster metadata to different formats with a single command.

Usage: [Export](../commands/export.md)

See documentation for all the [Commands](../commands/README.md)

 
 