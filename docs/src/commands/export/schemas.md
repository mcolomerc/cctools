# Export Schemas

```sh:no-line-numbers
cctools export schemas --help
```

```sh:no-line-numbers
 Command to export Schemas information.

Usage:
  cctcctools export schemas [flags]

Aliases:
  schemas, schemas-info, schemas-exp, schema

Flags:
  -h, --help   help for schemas

Global Flags:
  -c, --config string   config file  
  -o, --output string   Output format. Possible values: json, yaml, hcl, cfk, clink
```

Output format:

* JSON: `cctools export schemas --output json --config config.yaml`
* YAML: `cctools export schemas --output yaml --config config.yaml`
* CFK(YML): `cctools export schemas --output cfk --config config.yaml`
* CLINK(SH): `cctools export schemas --output clink --config config.yaml`
* HCL(TFVARS): `cctools export schemas --output hcl --config config.yaml`

## Configuration

Configure Subject export: `all` subject versions or only the `latest` version.

```yaml
export: 
  schemas: 
    version: latest  # default: all 
    subjects:
      version: latest # default: all 
```