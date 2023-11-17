# Import Topics

Import from JSON files all the Topics metadata to a Destination Cluster.

See [Topic Import](../commands/import/topics.md) documentation for more details.

Create a `config.yaml` file:

```sh:no-line-numbers
touch config.yaml
```

- Add the Destination Cluster connection configuration, the `destination` tag requires:
  
- Cluster: `kafka` section contains the source Kafka cluster connection configuration.
  - `bootstrapServer`: Source Cluster bootstrap server.
  - `clientProps`: Kafka client properties map.

Example for SASL_SSL (*config.yaml*):

```yaml
destination: 
  kafka:
    bootstrapServer: <bootstrap_server>
    clientProps:  
      - ssl.ca.location: "<path>/cacerts.pem" 
      - sasl.mechanisms: PLAIN
      - security.protocol: SASL_SSL
      - sasl.username: <admin_user>
      - sasl.password: <admin_password>
```

- Add the output path for the JSON files. 
  Using exported files. All the Topics metadata will be exported to `<export.output.path>/topics/json`, one file per Topic.
  Import uses same `export.output.path` as default.

```yaml
import:
  source: <path>
```

Topic JSON file example:

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

- Run the command:

```sh:no-line-numbers
cctools import topics --config config.yaml 
```

- Check the Destination cluster topic list 
  
For Confluent Cloud: 

```sh:no-line-numbers
confluent kafka topic list 
```
