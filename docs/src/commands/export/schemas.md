# Export Schemas

```sh:no-line-numbers
cctools export schemas --help
```

```sh:no-line-numbers
 Command to export Schemas information.

Usage:
  cctcctools export schemas [flags] 

Flags:
  -h, --help   help for schemas

Global Flags:
  -c, --config string   config file  
  -o, --output string   Output format. Possible values: json, yaml, hcl, cfk, clink
```

Output format:

* JSON: `cctools export schemas --output json --config config.yaml`
* YAML: `cctools export schemas --output yaml --config config.yaml`
* CFK(YML): `cctools export schemas --output cfk --config config.yaml`
* Excel(XLS): `cctools export schemas --output excel --config config.yaml`
 

## Configuration

Configure Subject export: `all` subject versions or only the `latest` version.

```yaml
export: 
  schemas: 
    version: latest  # default: all 
    subjects:
      version: latest # default: all 
```

## Export format

### JSON

Schemas:

Output path: `<export.output.path>/schemas/json`

```json
{
 "subject": "customer-value",
 "version": 1,
 "id": 100011,
 "schema": "{\"type\":\"record\",\"name\":\"Customer\",\"fields\":..."
}
```

Subjects:

`<export.output.path>/subjects/json`

```json
{
 "subject": "customer-value",
 "version": 1,
 "id": 100011,
 "schema": "{\"type\":\"record\",\"name\":\"Customer\",\"fields\":[..."
}
```

### YAML

Schemas:

`<export.output.path>/schemas/yaml`

```yaml
subject: payment-value
version: 1
id: 100064
schemaType: ""
schema: '{"type":"record","name":"Payment","namespace":"io.confluent.examples.clients.basicavro","fields":[{...'
```

Subjects:

`<export.output.path>/subjects/yaml`

```yaml
subject: payment-value
version: 1
id: 100064
schemaType: ""
schema: '{"type":"record","name":"Payment","namespace":"io.confluent.examples.clients.basicavro","fields":[{"name":"id","type":"string"},{"name":"amount","type":"double"},{"name":"email","type":"string"}]}'
```

### CFK 

CFK export configuration: 

```yaml
export: 
  schemas: 
    version: latest  # default: all 
    subjects:
      version: latest # default: all 
  cfk:
    namespace: confluent  
    kafkarestclass: kafka 
    schemaRegistryClusterRef: schemaregistry 
```

Output: `<export.output.path>/schemas/cfk`

Schema:

```yaml
apiVersion: platform.confluent.io/v1beta1
kind: Schema
metadata:
  name: schema_name-value
  namespace: confluent
spec:
  data:
    configRef: schema_name-value-config
    format: avro
  schemaRegistryClusterRef:
    name: schemaregistry
```

Config Map:
  
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: schema_name-value-config
  namespace: confluent
data:
  schema: '{"type":"record","name":"record","namespace":"org.apache.flink.avro.generated","fields":...'
```