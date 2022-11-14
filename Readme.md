
# Confluent Migration Tools

`cctools` is command Line tool for helping on migrations to Confluent Cloud or Confluent Platform.

This CLI uses Kafka REST API to extract and export all the resources from the Source cluster in order to replicate them on the target cluster. It was tested with Confluent Platform and Confluent Cloud clusters. 

It allows to export resources into different formats, that could be used as input for different tools like Confluent Cloud, Terraform, Confluent For Kubernetes or any other tool. 

<img src="./docs/image.png" width="500">


* [Installation](#installation)
  
* [Configuration](#configuration)
 
  * [Commands](#commands)
 
  * [Connection and Credentials](#connection)

  * [Output](#output)

  * [Resources](#resources)

  * [Exporters](#exporters)
  

## Installation

Go to [Releases](https://github.com/mcolomerc/cctools/releases) and Download your OS distribution.

## Configuration

The tool needs a configuration file (yml).

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
```

### Connection

* Rest Api URL: ```endpointUrl: <REST_ENDPOINT>```

* Credentials: 
  
  * ```key: <USER>``` or Confluent Cloud API_KEY.
  
  * ```secret: <PASSWORD>``` or Confluent Cloud API_SECRET  

```yaml
cluster: <CLUSTER_ID>
#bootstrap server
bootstrapServer: <BOOTSTRAP_SERVER> 
#REST endpint 
endpointUrl: <REST_ENDPOINT>
#Credentials
credentials: 
  key: <USER> # or CCloud API_KEY 
  secret: <PASSWORD> # or CCloud API_SECRET   
```

If certiticates are needed:

* Certificates:
   
  * `certFile`: Certificate file path
  
  * `keyFile`: Key file path
  
  * `CAFile`: CA file path


```yaml
cluster: <CLUSTER_ID>
#bootstrap server
bootstrapServer: <BOOTSTRAP_SERVER> 
#REST endpint 
endpointUrl: <REST_ENDPOINT>
#Credentials
credentials: 
  key: <USER>  
  secret: <PASSWORD>  
  # Certificates - Confluent Platform 
  certificates: 
    certFile: <CERT file path>  
    keyFile: <KEY file path>  
    CAFile: <CA File path>
```

#### **Schema Registry** 

Schema Registry connection configuration:

```yaml
#Schema Registry 
schemaRegistry: 
  endpointUrl: <Schema_Registry_Url>
#Credentials
  credentials: 
    key: <USER> # or CCloud API_KEY 
    secret: <PASSWORD> # or CCloud API_SECRET   
```

If certiticates are needed:

* Certificates:

  * `certFile`: Certificate file path
  
  * `keyFile`: Key file path
  
  * `CAFile`: CA file path

```yaml
#Schema Registry 
schemaRegistry: 
  endpointUrl: <Schema_Registry_Url>
#Credentials
  credentials: 
    key: <USER>  
    secret: <PASSWORD>  
  certificates: 
    certFile: <CERT file path>  
    keyFile: <KEY file path>  
    CAFile: <CA File path>
```

---

## Commands

`export`

Configuration:

* Resources to export, a list of resources to export, available values: `topics`, `consumer_groups`, `schemas`.  
* Output path
* Exporter configuration 
  * Specific configuration for each exporter (See Exporters)
* List of Exporters 
  * List of export formats: 
    * `json`: Json files
    * `yaml`: YAML files
    * `excel`: Excel files
    * `clink`: Cluster Linking commands (.sh and configuration files) 
    * `cfk`: Confluent For Kubernetes Custom Resources definitions. (YAML for Kubernetes)

```yaml
export:
  resources:
    - topics
    - consumer_groups
    - schemas
  topics:
    exclude: _confluent
  exporters: 
  - excel
  - yaml 
  - json  
  output: output #Output Path
``` 

---

### Resources

Required: Configure resources to export.

* Export Topics including Topic configuration: ```topics```

```yaml
export:  
  resources: 
    - topics
```

See [Topics](docs/Topics.md).

* Export Consumer Groups information: ```consumer_groups```

```yaml
export:  
  resources: 
    - consumer_groups
```

* Export Schema Registry information: ```schemas```

```yaml
export:  
  resources: 
    - schemas
```

See [Schemas](docs/Schemas.md)

---

### Output

Configure the output folder, it will be created if it does not exist.

Example: All the export files will be stored into the ```output``` folder (it will be created if necessary).
  
```yaml
export: 
  output: output 
```

1. Each `resource` will create a folder inside the `output` target.

2. Each exporter will create a folder inside the `resource` folder.

**Example**: Exporting Topics to JSON will generate: `output/topics/json/topics.json`

---

### **Exporters** 

```cctools export``` supports different exporters by configuration: `json`, `yaml`,`excel`, `clink`, `cfk`

* JSON: `json`
* YAML: `yaml`
* Excel: `excel`
* [CLinkExporter](docs/CLinkExporter.md): `clink`
* [CFKExporter](docs/CFKExporter.md): `cfk`

Example: Use the following configuration to export to *YAML* format only:

```yaml
export:
  exporters:  
  - yaml  
```

Example: Use the following configuration to export to *Excel*, YAML and JSON formats:

```yaml
export:
  exporters: 
  - excel
  - yaml 
  - json  
```

---

# Sources

## Execute

`go run main.go  export  --config config_cloud.yml`

## DEBUG

Enable Debug mode:`LOG=DEBUG` for extra logging.

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



	https://docs.confluent.io/cloud/current/connectors/index.html#supported-connectors
Managed Cloud connectors
Networking 
https://docs.confluent.io/cloud/current/connectors/index.html#networking
Limitations
https://docs.confluent.io/cloud/current/connectors/limits.html#limitations 
            Own Managed Connectors
Connect Kafka Connect to Confluent Cloud
https://docs.confluent.io/cloud/current/cp-component/connect-cloud-config.html
