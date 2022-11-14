# Schemas

Configure Subject export: `all` subject versions or only the `latest` version.

```yaml
export:
  resources:  
    - schemas 
  schemas: 
    version: latest  # default: all 
    subjects:
      version: latest # default: all 
```