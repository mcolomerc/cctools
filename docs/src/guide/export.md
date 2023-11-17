# Export Topics

Export to JSON files all the Topics metadata from a Source Cluster. See [Topic Exporters](../commands/export/topics.md) documentation for using other exporters.

Create a `config.yaml` file:

```sh:no-line-numbers
touch config.yaml
```

- Add the Source Cluster connection configuration, the `source` tag requires:
  
- Cluster: `kafka` section contains the source Kafka cluster connection configuration.
  - `bootstrapServer`: Source Cluster bootstrap server.
  - `clientProps`: Kafka client properties map.

Example for SASL_SSL (*config.yaml*):

```yaml
source: 
  kafka:
    bootstrapServer: <bootstrap_server>
    clientProps:  
      - ssl.ca.location: "<path>/cacerts.pem" 
      - sasl.mechanisms: PLAIN
      - security.protocol: SASL_SSL
      - sasl.username: <admin_user>
      - sasl.password: <admin_password>
```

- Add the output path for the JSON files, all the Topics metadata will be exported to `<export.output.path>/topics/json`, one file per Topic.

```yaml
export:
  output: <path>
```

- Filtering Topics. Exclude.

```yaml
export:
  output: <path>
  topics:
    exclude: <topic_name_substring>  #_confluent
```

- Run the command:

```sh:no-line-numbers
cctools export topics --config config.yaml --output json
```

- Output folder. One JSON file per Topic exported.

JSON Topic example:

```json
{
 "name": "demo.topic",
 "partitions": 4,
 "replicationFactor": 3,
 "minIsr": "2",
 "retentionTime": "604800000",
 "configs": [
  {
   "name": "confluent.stray.log.max.deletions.per.run",
   "value": "72"
  },
  {
   "name": "confluent.stray.log.delete.delay.ms",
   "value": "604800000"
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
```