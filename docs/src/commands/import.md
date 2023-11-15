---
prev: ./export.md
next: ./copy.md
---

# Import

Import metadata from source files and create resources into Destination cluster

```sh:no-line-numbers
cctools import --help
```

Usage:

```sh:no-line-numbers
Command to import cluster resources  to another cluster.

Usage:
  cctools import [flags]
  cctools import [command]

Aliases:
  import, i

Available Commands:
  topics      Import Topics Info

Flags:
  -h, --help   help for import

Global Flags:
  -c, --config string   config file
```

## Configuration

Source path:

```yaml
import:
  source: <path>
```

::: info
The source path is `export.output` folder by default.
:::
