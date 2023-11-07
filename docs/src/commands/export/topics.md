# Export Topics

Export Topics metadata to different formats.

Includes Topic configuration and Topic ACLs.

```sh:no-line-numbers
cctools export topics --help`
```

```sh:no-line-numbers
 Command to export Topics information.

Usage:
  cctools export topics [flags]

Aliases:
  topics, topic-info, topic-exp, tpc

Flags:
  -h, --help   help for topics

Global Flags:
  -c, --config string   config file 
  -o, --output string   Output format. Possible values: json, yaml, hcl, cfk, clink
```

## Configuration

* Using Topic Exporter Configuration to exclude some topics.

All topics names containing `_confluent` will be excluded.

```yaml
export: 
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

## Export format

Required `--output`

Output format:

* JSON:

```sh
  cctools export topics --output json --config config.yaml
```

* YAML:
  
```sh:no-line-numbers
  cctools export topics --output yaml --config config.yaml
```

* CFK(YML):
  
```sh:no-line-numbers
  cctools export topics --output cfk --config config.yaml
```
  
* CLINK(SH):

```sh:no-line-numbers
cctools export topics --output clink --config config.yaml
```

* HCL(TFVARS):

```sh:no-line-numbers
cctools export topics --output hcl --config config.yaml
```

## Example

```sh:no-line-numbers
cctools export topics --output json --config config.yaml
```

* Source cluster configuration (config.yaml)
* Exclude consiguration.

Exporter will create a `JSON` file per topic selected.

Topic selected : `demo.topic`. `demo.topic.json` file under `output/topics/json` folder.

```json
{
 "name": "demo.topic",
 "partitions": 4,
 "replicationFactor": 3,
 "minIsr": "2",
 "retentionTime": "604800000",
 "configs": [
  {
   "name": "index.interval.bytes",
   "value": "4096"
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
}
```

### YAML

OutPut: `<output_path> / <cluster_ID>_<resource>.yaml`

Output Sample for `topics` resource:

```yaml
- name: demo.topic
  partitions: 1
  replicationfactor: 3
  configs:
  - name: cleanup.policy
    value: compact
  - name: compression.type
    value: producer
  ...
```

### CFK

From the selected resources export to [Confluent For Kubernetes (CFK) Custom Resources](https://docs.confluent.io/operator/current/co-manage-topics.html#create-ak-topic) `KafkaTopic`.

Output: Exporter will create a `<output_path>/<clusterid>_<resource>_<topicName>.yml` file for each Topic.

Example:

```yml
apiVersion: platform.confluent.io/v1beta1
kind: KafkaTopic
metadata:
  name: demo.topic
  namespace: confluent
spec:
  replicas: 3
  partitionCount: 6
  configs:
    cleanup.policy: delete
    compression.type: producer
    ...
  kafkaRestClassRef:
    name: kafka

```

### Cluster Link

From the selected `topics` export [Confluent Cloud Cluster Link](https://docs.confluent.io/cloud/current/multi-cloud/overview.html) scripts and configuration.

Output: The export will generate:

* Cluster Link creation script (.sh), including topic mirrors from selected `topics` if `autocreate` is `false`
* Topic promotion script (.sh)
* Clean up script (.sh)
* Cluster Link configuration file (.properties), including `auto.create.mirror.topics.filters` from selected `topics`

### HCL

Exporting Topics to HCL will generate: `output/topics/tfvars/topics.tfvars`.

```json
environment = "<ENV_ID>"

cluster = "<CLUSTER_ID>"

rbac_enabled = false

serv_account = {
  name = "<SERVICE_ACCOUNT>"
  role = "CloudClusterAdmin"
}
topics = [{
  name       = "orders"
  partitions = 12
  config = {
    "cleanup.policy"                          = "delete"
    "compression.type"                        = "producer"
    ...
```

The output could be used with [Terraform Topics Module](https://github.com/mcolomerc/terraform-confluent-topics) to create topics on Confluent cloud destination cluster.

### Excel

Output: `<output_path>/<cluster_ID>_<resource>.xlsx`