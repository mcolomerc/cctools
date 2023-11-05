# Configuration

`cctools` commands require a configuration file using `--config` (yml) with the source cluster connection.

Create a `config.yml` file.

## Source Cluster

Required. 

`source` section contains the source cluster connection configuration.
`clientProps` contains the Kafka client properties map.

Example with `SASL_SSL`

```yaml
source: 
  bootstrapServer: <bootstrap_server>
  clientProps:  
    - ssl.ca.location: "<path>/cacerts.pem" 
    - sasl.mechanisms: PLAIN
    - security.protocol: SASL_SSL
    - sasl.username: <username>
    - sasl.password: <password>
```

## Destination Cluster

Optional. Some commands require a destination cluster connection configuration.

```yaml
destination: 
  bootstrapServer: <bootstrap_server>.confluent.cloud:9092
  clientProps:
    - sasl.mechanisms: PLAIN
    - security.protocol: SASL_SSL
    - sasl.username: <API_KEY>
    - sasl.password: <API_SECRET>
```

## Schema Registry

Required. Source Schema Registry connection configuration:

```yaml
#Schema Registry 
schemaRegistry: 
  endpointUrl: <Schema_Registry_Url>
#Credentials
  credentials: 
    key: <USER> # or CCloud API_KEY 
    secret: <PASSWORD> # or CCloud API_SECRET   
```

## Exporters configuration

* `output` path.
* Exporter configuration:
  * Specific configuration for each exporter (See Exporters)
  * `exclude` resources

### Topics

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

## Schemas

Configure Subject export: `all` subject versions or only the `latest` version.

```yaml
export: 
  schemas: 
    version: latest  # default: all 
    subjects:
      version: latest # default: all 
```

## External resources

Add external Git repositories to the `output`.

Provide a map as `target_dir`: `git url`.

The repository will be cloned into `output/target_dir`

```yaml
export:
  output: output 
  git:
    scripts: https://github.com/mddunn/ccloud-migration-scripts
    terraform: https://github.com/mcolomerc/terraform-confluent-topics
```

In the example above:  

* The `https://github.com/mddunn/ccloud-migration-scripts` repository will be cloned into `output/scripts`

* The `https://github.com/mcolomerc/terraform-confluent-topics` repository will be cloned into `output/terraform`
