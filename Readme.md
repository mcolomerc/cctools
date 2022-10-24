
# Confluent Cloud Tools 


Configuration file: ```--config config.yml```

```yaml
environment: <CCLOUD_ENVIRONMENT_ID>
cluster: <CCLOUD_CLUSTER_ID>
bootstrapServer: <CCLOUD_BOOTSTRAP_SERVER>    
ccloudUrl: <CCLOUD_CLUSTER_REST_URL>
apiKey: <CCLOUD_API_KEY>
apiSecret: <CCLOUD_API_SECRET>
```

## Commands

### Export 

```cctools export --config config.yml```

OutPut: (<cluster_ID>_Topics.xlsx)

Output Sample: 

| Topic	| Partitions |	Replication Factor | Configs |
|-------|------------|---------------------|---------|
|_confluent-command |	1 |	3 |	cleanup.policy=compact compression.type=producer delete.retention.ms=86400000 ...|
| my-topic | 6 | 3 | cleanup.policy=delete compression.type=producer delete.retention.ms=86400000 ...| 
 
... 


# Sources

## Releaser

https://goreleaser.com/install/

```brew install goreleaser```

```goreleaser build --snapshot --rm-dist```

### Binary

Mac/OS:

```./dist/cctools_darwin_amd64_v1/cctools export --config config.yml```
