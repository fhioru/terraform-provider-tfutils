# Terraform provider for formatting JSON strings

This plugin will format / beautify a JSON string.

## Usage

```hcl
data "json-formatter_format" "test" {
    json = "{\"a\":\"b\",\"myObj\":{\"prop1\":1}}"
    indent = "    "
}

output "test_output" {
    value = data.json-formatter_format.test.result
}
```

This will output:

```
Outputs:

test_output = <<EOT
{
    "a": "b",
    "myObj": {
        "prop1": 1
    }
}
EOT
```

### Arguments
|Name|Required|Description|
|---|---|---|
|`json`|Yes|A JSON string.|
|`indent`|No|The character(s) to use for an indent. Has a default value of two spaces.|