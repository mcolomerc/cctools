#### Confluent For Kubernetes 

Configuration requires:

* `namespace`: target namespace `string`
* `kafkarestaclass`: Kafka Rest Class name `string`

```yaml
export:
  cfk:
    namespace: confluent  
    kafkarestclass: kafka 
  exporters:
  - cfk  
```