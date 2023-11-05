#### Cluster Link

Configuration requires:

 * `name`: Cluster Link name `string`
 * `destination`: Destination cluster ID `string`
 * `autocreate`: Autocreate topics `true|false`
 * `sync`: 
    * `offset`: Offset sync `true|false`
    * `acl`: Acl Sync `true|false` 


```yaml
export:
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
