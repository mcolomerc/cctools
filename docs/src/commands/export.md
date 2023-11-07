---
prev: ./copy.md
next: ./import.md
---

# Export

```sh:no-line-numbers
cctools export --help
```

```sh:no-line-numbers
Command to export cluster information.

Usage:
  cctools export [flags]
  cctools export [command]

Aliases:
  export, export-info, cluster-export, confluent-exp, exp

Available Commands:
  schemas     Export Schemas Info
  topics      Export Topics Info

Flags:
  -h, --help            help for export
  -o, --output string   Output format. Possible values: json, yaml, hcl, cfk, clink

Global Flags:
  -c, --config string   config file 

Use "cctools export [command] --help" for more information about a command.
```

## Output folder

Configure the output folder, it will be created if it does not exist.

Example: All the export files will be stored into the ```output``` folder (it will be created if necessary).
  
```yaml
export: 
  output: output 
```

1. Each `resource` will create a folder inside the `output` target.

2. Each exporter will create a folder inside the `resource` folder.

**Example**: Exporting Topics to JSON will generate: `output/topics/json/topics.json`

## Resources

- [Topics](/commands/export/topics.md)

- [Consumer Groups](/commands/export/consumer-groups.md)

- [Schemas](/commands/export/schemas.md)

## Exporters

`--output` flag is required.

- [JSON](/commands/export/topics.md)
- [YAML](/commands/export/topics.md)
- [HCL](/commands/export/topics.md)
- [Confluent Cloud](/commands/export/topics.md)
- [Confluent Cloud Link](/commands/export/topics.md)

