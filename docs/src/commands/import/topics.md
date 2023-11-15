--- 
next: ./schemas.md
---
# Import Topics

Import Topics metadata from source files and create destination Topics into Destination cluster

Includes Topic configuration and Topic ACLs.

```sh:no-line-numbers
cctools import topics --help`
```

```sh:no-line-numbers
Command to import from source files and create destination Topics.

Usage:
  cctools import topics [flags] 

Flags:
  -h, --help   help for topics

Global Flags:
  -c, --config string   config file
```

::: tip
Works with exported JSON files. See [Export Topics](../export/topics.md) for more information.
:::

## Configuration

Destination cluster.

```yaml
destination: 
  kafka:
    bootstrapServer: <bootstrap_server>.confluent.cloud:9092
    clientProps:
      - sasl.mechanisms: PLAIN
      - security.protocol: SASL_SSL
      - sasl.username: <API_KEY>
      - sasl.password: <API_SECRET>
```

See [Configuration](../config/README.md)