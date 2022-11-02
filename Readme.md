
# Confluent Migration Tools

Command Line Tools for helping on migrations to Confluent Cloud or Confluent Platform.
 

## Configuration

Configuration file: ```--config config.yml```

```yaml 
#cluster id
cluster: <CLUSTER_ID>
#bootstrap server
bootstrapServer: <BOOTSTRAP_SERVER> 
#REST endpint 
endpointUrl: <REST_ENDPOINT>
#Credentials
credentials: 
  key: <USER> # or CCloud API_KEY 
  secret: <PASSWORD> # or CCloud API_SECRET  
  # Certificates - Confluent Platform 
  certificates: 
    certFile: <CERT file path>  
    keyFile: <KEY file path>  
    CAFile: <CA File path>
ccloud:
  environment: <ENVIRONMENT_ID>  
export:
  topics:
    exclude: _confluent
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

#### Output

Configure the output folder, it will be created if it does not exist. 

Example: All the export files will be stored into the ```output``` folder (it will be created if necessary).
  
```yaml
export: 
  output: output 
```

#### Exclude 

##### Topics

Exclude Topics by name containing ```string```.

```yaml
export:
  topics:
    exclude: _confluent
```

#### Include

##### Topics

Include only Topics by name containing ```string```.

```yaml
export:
  topics: 
    include: _confluent
```

#### Exclude & Include

```include``` can be used with ```exclude``` rules.  

Example: Exclude all ```_confluent``` topics but include ```_confluent_balancer``` topics.

```yaml
export:
  topics:
    exclude: _confluent
    include: _confluent_balancer
```

Export result:  

* *_confluent_balancer_api_state*
* *_confluent_balancer_broker_samples*
* *_confluent_balancer_partition_samples*
* *my-topic*

Consider that ```exclude``` will be applied first, so with the following configuration, ```exclude``` will not take effect since all the ```_confluent``` topics will be included by the ```include```.

```yaml
export:
  topics:
    exclude: _confluent_balancer
    include: _confluent
```

---

#### Exporters

```cctools export``` supports different exporters by configuration.

Example: Use the following configuration to export to YAML format only:

```yaml
export:
  exporters:  
  - yaml  
```

Example: Use the following configuration to export to Excel, YAML and JSON formats: 

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
| ... | | | | 

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

---

# Sources

## Releaser

https://goreleaser.com/install/

```brew install goreleaser```

```goreleaser build --snapshot --rm-dist```

### Binary

Mac/OS:

```./dist/cctools_darwin_amd64_v1/cctools export --config config.yml```

 