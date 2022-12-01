# Confluent For Kubernetes

Configuration requires:

* `namespace`: target namespace `string`
* `kafkarestaclass`: Kafka Rest Class name `string`
* `schemaRegistryClusterRef`: Schema Registry cluster ref. `string` for Schema exporter.

```yaml
export:
  cfk:
    namespace: confluent  
    kafkarestclass: kafka 
    schemaRegistryClusterRef: schemaregistry
  exporters:
  - cfk  
```
