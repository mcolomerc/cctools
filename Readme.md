
# Confluent Migration Tools

`cctools` is command Line tool for helping on migrations to Confluent Cloud or Confluent Platform.

This CLI uses Kafka REST API to extract and export all the resources from the Source cluster in order to replicate them on the target cluster. It was tested with Confluent Platform and Confluent Cloud clusters. 

It allows to export resources into different formats, that could be used as input for different tools like Confluent Cloud, Terraform, Confluent For Kubernetes or any other tool.

<img src="./docs/image.png" width="500">

## Install

Go to [Releases](https://github.com/mcolomerc/cctools/releases) and Download your OS distribution.

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
  resources:
    - topics
    - consumer_groups
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

####  Resources

Required: Configure resources to export.

* Export Topics including Topic configuration: ```topics```

```yaml
export:  
  resources: 
    - topics
```

* Export Consumer Groups information: ```consumer_groups```

```yaml
export:  
  resources: 
    - consumer_groups
```

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

OutPut: (<output_path>/<cluster_ID>_<resource>.xlsx)

Output Sample for ```topics``` resource:

| Topic	| Partitions |	Replication Factor | Configs |
|-------|------------|---------------------|---------|
|_confluent-command |	1 |	3 |	cleanup.policy=compact compression.type=producer delete.retention.ms=86400000 ...|
| my-topic | 6 | 3 | cleanup.policy=delete compression.type=producer delete.retention.ms=86400000 ...| 
| ... | | | | 

##### ```json```

OutPut: (<output_path>/<cluster_ID>_<resource>.json)

Output Sample for ```topics``` resource:

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

OutPut: (<output_path>/<cluster_ID>_<resource>.yaml)

Output Sample for ```topics``` resource:

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

##### Confluent For Kubernetes `cfk`

From the selected `topics` export to [Confluent For Kubernetes (CFK) Custom Resources](https://docs.confluent.io/operator/current/co-manage-topics.html#create-ak-topic) `KafkaTopic`. 

Configuration requires:

* namespace
* kafkarestaclass

Example:

```yaml
export:
  resources: 
    - topics
  topics:
    exclude: connect
  cfk:
    namespace: confluent  
    kafkarestclass: kafka 
  exporters:  
  - cfk
```

Output: Exporter will create a <output_path>/<clusterid>_topic_<topicName>.yml file for each Topic. 

Example:

```yml
apiVersion: platform.confluent.io/v1beta1
kind: KafkaTopic
metadata:
  name: user_transactions
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

##### Cluster Link ```clink```

From the selected `topics` export [Confluent Cloud Cluster Link](https://docs.confluent.io/cloud/current/multi-cloud/overview.html) scripts and configuration.

Output: The export will generate:

* Cluster Link creation script (.sh), including topic mirrors from selected `topics` if `autocreate` is `false`
* Topic promotion script (.sh)
* Clean up script (.sh)
* Cluster Link configuration file (.properties), including `auto.create.mirror.topics.filters` from selected `topics`

Exporter configuration:

```yaml
export: 
  resources: 
    - topics
  topics:
    exclude: connect
  clink:
    name: <CLUSTER_LINK_NAME>
    destination: <DESTINATION_CLUSTER_ID>
    autocreate: true | false
    sync: 
      offset: true | false
      acl: true | false 
  exporters:  
  - clink 
``` 

---

# Sources

## Releaser

https://goreleaser.com/install/

```brew install goreleaser```

```goreleaser build --snapshot --rm-dist```

### Binary

MacOS:

```./dist/cctools_darwin_amd64_v1/cctools export --config config.yml```

### CI/CD

 There are 2 `github actions` get on the repo:

 1. `pr-tag`: Create a tag from every PR on the repo. You need to specify #major/#minor/#patch on the cluster for better version control. If not minor version will be created

 2. `release`: Create a new release from the TAG created by the previous tag. This action in created on top of `goreleaser` and will create binaries for all the common distributions. 
