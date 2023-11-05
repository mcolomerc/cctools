# Copy Topics

Copy Topics metadata to destination cluster.

Includes Topic configuration and Topic ACLs.

```sh
cctools copy topics --config config.yml
```

Usage:

```sh
Command to copy from source Kafka and create destination Topics.

Usage:
  cctools copy topics [flags]

Aliases:
  topics, topic-cp, tpic-cp, tpc

Flags:
  -h, --help   help for topics

Global Flags:
  -c, --config string   config file  
```

## Configuration

* Using Topic copyer Configuration to exclude some topics.

All topics names containing `_confluent` will be excluded.

```yaml
copy: 
  topics:
    exclude: _confluent 
```

* Topic ACLs - Principals Mapping

All the Topic ACLs where `principal: User:test` will be created as `principal: User:sa-xyroox` on the Destination.

```yaml
principals:
  - "test": "sa-xyroox"
```

[Configuration](../config/README.md)
