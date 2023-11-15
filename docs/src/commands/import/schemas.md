---
prev: ./topics.md 
---
# Import Schemas

Import Schemas from source files and create destination Subjects into Destination Schema Registry.

Subjects are created on the destination keeping the schema `id` and the schema `version`.

:::warning
If the schema `id` already exists on the destination, Schema Registry call will fail with an error when switching to `IMPORT`.
:::

```sh:no-line-numbers
cctools import schemas --help`
```

```sh:no-line-numbers
Command to import from source files and create destination Schemas.

Usage:
  cctools import schemas [flags]

Flags:
  -h, --help   help for schemas

Global Flags:
  -c, --config string   config file
```

::: tip
Works with JSON files. See [Export Schemas](../export/schemas.md) for more information.
:::

Example input JSON file:

```json
{
    "schema":"{\"type\":\"record\",\"name\":\"value_a1\",\"namespace\":\"com.mycorp.mynamespace\",\"fields\":[{\"name\":\"field1\",\"type\":\"string\"}]}",
    "version":19,
    "id":101010,
    "schemaType": "AVRO",
    "subject": "my-value" 
}
```

## Configuration

Destination cluster.

```yaml
destination: 
  schemaRegistry:
    endpointUrl: <SCHEMA_REGISTRY_URL>
    credentials:
      key: <USER>
      secret: <PASSWORD>
```

See [Configuration](../config/README.md)
