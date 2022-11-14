# Extract Topic Information 

Export Topics including Topic configuration: ```topics```

```yaml
export:
  resources:
    - topics
```

## Filtering

**Exclude** 

Exclude Topics by name containing ```string```.

```yaml
export:
  topics:
    exclude: _confluent
```

**Include** 

Include only Topics by name containing ```string```.

```yaml
export:
  topics: 
    include: _confluent
```

**Exclude & Include**

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

## Exporters

**```json```**

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

**```excel```**

OutPut: (<output_path>/<cluster_ID>_<resource>.xlsx)

Output Sample for ```topics``` resource:

| Topic	| Partitions |	Replication Factor | Configs |
|-------|------------|---------------------|---------|
|_confluent-command |	1 |	3 |	cleanup.policy=compact compression.type=producer delete.retention.ms=86400000 ...|
| my-topic | 6 | 3 | cleanup.policy=delete compression.type=producer delete.retention.ms=86400000 ...| 
| ... | | | | 



**```yaml```**

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

**`cfk`**

From the selected resources export to [Confluent For Kubernetes (CFK) Custom Resources](https://docs.confluent.io/operator/current/co-manage-topics.html#create-ak-topic) `KafkaTopic`. 

Output: Exporter will create a <output_path>/<clusterid>_<resource>_<topicName>.yml file for each Topic. 

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

**```clink```**

From the selected `topics` export [Confluent Cloud Cluster Link](https://docs.confluent.io/cloud/current/multi-cloud/overview.html) scripts and configuration.

Output: The export will generate:

* Cluster Link creation script (.sh), including topic mirrors from selected `topics` if `autocreate` is `false`
* Topic promotion script (.sh)
* Clean up script (.sh)
* Cluster Link configuration file (.properties), including `auto.create.mirror.topics.filters` from selected `topics`

 