# Consumer Groups

Consumer groups are a way of grouping Kafka consumers together to consume a topic. Each consumer in a group will consume from a unique subset of partitions in the topic. This allows you to horizontally scale your consumers while still maintaining the ordering guarantees of a single partition.

Exporting consumer groups will generate a file per consumer group.

```sh:no-line-numbers
cctools export consumer-groups --help
```

Usage:

```sh:no-line-numbers
 Command to export Consumer Group information.

Usage:
  cctools export consumer-groups [flags]

Aliases:
  consumer-groups, cg, cg-info, cg-exp, cgroup

Flags:
  -h, --help   help for consumer-groups

Global Flags:
  -c, --config string   config file
  -o, --output string   Output format. Possible values: json, yaml, hcl, cfk, clink
```
