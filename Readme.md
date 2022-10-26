
# Confluent Cloud Tools

Command Line Tools for Confluent, provides extra features not included in [Confluent CLI](https://docs.confluent.io/confluent-cli/current/overview.html).

## Configuration

Configuration file: ```--config config.yml```

```yaml
environment: <CCLOUD_ENVIRONMENT_ID>
cluster: <CCLOUD_CLUSTER_ID>
bootstrapServer: <CCLOUD_BOOTSTRAP_SERVER>    
ccloudUrl: <CCLOUD_CLUSTER_REST_URL>
apiKey: <CCLOUD_API_KEY>
apiSecret: <CCLOUD_API_SECRET>
export:
  exporters: 
  - excel
  - yaml 
  - json  
  output: output #Output Path
```

## Commands

### Export

Export Topic information:

```cctools export --config config.yml```

#### Exporters

```cctools export``` supports different exporters by configuration.

Configure the output folder, it will be created if it does not exist. 

- Example: All the export files will be stored into the ```output``` folder (it will be created if necessary).
  
```yaml
export: 
  output: output 
```

- Example: Use the following configuration to export to YAML format only:

```yaml
export:
  exporters:  
  - yaml  
```

- Example: Use the following configuration to export to Excel, YAML and JSON formats: 

```yaml
export:
  exporters: 
  - excel
  - yaml 
  - json  
```

##### ```excel```

OutPut: (<output_path>/<cluster_ID>_Topics.xlsx)

Output Sample: 

| Topic	| Partitions |	Replication Factor | Configs |
|-------|------------|---------------------|---------|
|_confluent-command |	1 |	3 |	cleanup.policy=compact compression.type=producer delete.retention.ms=86400000 ...|
| my-topic | 6 | 3 | cleanup.policy=delete compression.type=producer delete.retention.ms=86400000 ...| 
 
... 

##### ```json```

OutPut: (<output_path>/<cluster_ID>_Topics.json)

Output Sample:

```json
[
 {
  "Name": "_confluent-command",
  "Partitions": 1,
  "ReplicationFactor": 3,
  "Configs": [
   {
    "Name": "cleanup.policy",
    "Value": "compact"
   },
   {
    "Name": "compression.type",
    "Value": "producer"
   },
   ...
```

##### ```yaml```

OutPut: (<output_path>/<cluster_ID>_Topics.yaml)

Output Sample: 

```yaml
- name: _confluent-command
  partitions: 1
  replicationfactor: 3
  configs:
  - name: cleanup.policy
    value: compact
  - name: compression.type
    value: producer
  ...
```

# Sources

## Releaser

https://goreleaser.com/install/

```brew install goreleaser```

```goreleaser build --snapshot --rm-dist```

### Binary

Mac/OS:

```./dist/cctools_darwin_amd64_v1/cctools export --config config.yml```

 